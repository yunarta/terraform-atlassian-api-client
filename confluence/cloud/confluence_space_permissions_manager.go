package cloud

import (
	"fmt"
	"github.com/yunarta/golang-quality-of-life-pack/collections"
	"github.com/yunarta/terraform-atlassian-api-client/confluence"
	"slices"
	"strings"
)

type PermissionActorType int
type PermissionsChangeType int

const (
	permissionActorUser PermissionActorType = iota
	permissionActorGroup
)

const (
	permissionAdding PermissionsChangeType = iota
	permissionRemoving
)

type SpacePermissionsManager struct {
	spaceKey string
	client   *ConfluenceClient

	objectPermissions        *confluence.ObjectPermissions
	changeRequests           map[string]*permissionChangeRequest // map[role.id]changeRequest
	permissionsMapForRemoval map[string]string
}

type permissionChangeRequest struct {
	addingUsers    []string
	removingUsers  []string
	addingGroups   []string
	removingGroups []string
}

// newPermissionChangeRequest initializes a new permissionChangeRequest.
func newPermissionChangeRequest() *permissionChangeRequest {
	return &permissionChangeRequest{
		addingUsers:    make([]string, 0),
		removingUsers:  make([]string, 0),
		addingGroups:   make([]string, 0),
		removingGroups: make([]string, 0),
	}
}

func (manager *SpacePermissionsManager) ReadPermissions() (*confluence.ObjectPermissions, error) {
	space, err := manager.client.SpaceService().Get(manager.spaceKey)
	if err != nil {
		return nil, err
	}

	permissions, err := manager.client.SpacePermissionsService().GetPermissions(space.Id)
	if err != nil {
		return nil, err
	}

	userIds, groupIds := manager.getUserGroupIds(permissions)
	manager.client.ActorLookupService().RegisterAccountIds(userIds...)
	manager.client.ActorLookupService().RegisterGroupIds(groupIds...)

	groupRoles := make(map[string]confluence.GroupPermissions)
	userRoles := make(map[string]confluence.UserPermissions)
	permissionsMapForRemoval := make(map[string]string)

	// now armed with actor lookup service, we can build the same ObjectPermissions just like we have for jira, bitbucket and bamboo

	for _, permission := range *permissions {
		key := fmt.Sprintf("%s:%s:%s", permission.Principal.Id, permission.Operation.Key, permission.Operation.Target)
		permissionsMapForRemoval[key] = permission.Id
		if permission.Principal.Type == confluence.PrincipalTypeUser {
			username := manager.client.ActorLookupService().FindUserById(permission.Principal.Id)
			manager.addUserRole(userRoles, permission.Operation, permission.Principal.Id, username)

		} else if permission.Principal.Type == confluence.PrincipalTypeGroup {
			groupName := manager.client.ActorLookupService().FindGroupById(permission.Principal.Id)
			manager.addGroupRole(groupRoles, permission.Operation, permission.Principal.Id, groupName)
		}
	}

	userRolesList := collections.GetValuesOfMap(userRoles)
	groupRolesList := collections.GetValuesOfMap(groupRoles)

	manager.permissionsMapForRemoval = permissionsMapForRemoval
	manager.objectPermissions = &confluence.ObjectPermissions{
		Users:  userRolesList,
		Groups: groupRolesList,
	}

	return manager.objectPermissions, nil
}

func (manager *SpacePermissionsManager) addGroupRole(groupRoles map[string]confluence.GroupPermissions, permission confluence.SpacePermissionOperation, principalId, name string) {
	// First we retrieve the group from the map if available
	group, exists := groupRoles[principalId]
	if !exists {
		// The group not in the map, so we add the group for the first time
		group = confluence.GroupPermissions{
			Name:        name,
			AccountId:   principalId,
			Permissions: []string{},
		}
	}
	group.Permissions = append(group.Permissions, permission.ToSlug())
	groupRoles[principalId] = group
}

func (manager *SpacePermissionsManager) addUserRole(userRoles map[string]confluence.UserPermissions, permission confluence.SpacePermissionOperation, principalId, name string) {
	// First we retrieve the user from the map if available
	user, exists := userRoles[principalId]
	if !exists {
		user = confluence.UserPermissions{
			Name:        name,
			AccountId:   principalId,
			Permissions: []string{},
		}
	}
	user.Permissions = append(user.Permissions, permission.ToSlug())
	userRoles[principalId] = user
}

func (manager *SpacePermissionsManager) getUserGroupIds(permissions *[]confluence.SpacePermission) ([]string, []string) {
	usersMap := make(map[string]bool)
	groupsMap := make(map[string]bool)

	for _, permission := range *permissions {
		if permission.Principal.Type == confluence.PrincipalTypeUser {
			usersMap[permission.Principal.Id] = true
		} else if permission.Principal.Type == confluence.PrincipalTypeGroup {
			groupsMap[permission.Principal.Id] = true
		}
	}

	return collections.GetKeysOfMap(usersMap), collections.GetKeysOfMap(groupsMap)
}

func (manager *SpacePermissionsManager) UpdateUserRoles(username string, newRoles []string) error {
	// read assigned roles for selected group
	accountId := manager.client.ActorLookupService().FindUser(username)
	user := manager.objectPermissions.FindUser(accountId)

	if user != nil {
		adding, removing := user.DeltaPermissions(newRoles)
		manager.prepareRoleChanges(permissionActorUser, accountId, adding, removing)
	} else if len(newRoles) > 0 {
		// If the item is not found but there are new permissions, add them.
		manager.makePermissionChange(permissionActorUser, accountId, permissionAdding, newRoles)
	}
	return nil
}

func (manager *SpacePermissionsManager) UpdateGroupRoles(groupName string, newRoles []string) error {
	// read assigned roles for selected group
	groupId := manager.client.ActorLookupService().FindGroup(groupName)
	group := manager.objectPermissions.FindGroup(groupId)

	if group != nil {
		adding, removing := group.DeltaPermissions(newRoles)
		manager.prepareRoleChanges(permissionActorGroup, groupId, adding, removing)
	} else if len(newRoles) > 0 {
		// If the item is not found but there are new permissions, add them.
		manager.makePermissionChange(permissionActorGroup, groupId, permissionAdding, newRoles)
	}
	return nil
}

func (manager *SpacePermissionsManager) prepareRoleChanges(permissionActor PermissionActorType, actorId string, adding, removing []string) {
	if len(adding) > 0 {
		// Add new permissions if there are any to add.
		manager.makePermissionChange(permissionActor, actorId, permissionAdding, adding)
	}
	if len(removing) > 0 {
		// Remove old permissions if there are any to remove.
		manager.makePermissionChange(permissionActor, actorId, permissionRemoving, removing)
	}
}

func (manager *SpacePermissionsManager) makePermissionChange(permissionActor PermissionActorType, actorId string, changeType PermissionsChangeType, permissions []string) {
	actions := map[PermissionActorType]map[PermissionsChangeType]func(string, string){
		permissionActorUser: {
			permissionAdding: func(permission, actorId string) {
				// permission is a string containing "key_target"
				// the key it self may contains "_" so we need to split at the last "_"
				permissionParts := strings.Split(permission, "_")
				_, _ = manager.client.SpacePermissionsService().AddPermission(manager.spaceKey, confluence.AddPermissionRequest{
					Subject: confluence.SpacePermissionSubject{
						Type: confluence.PrincipalTypeUser,
						Id:   actorId,
					},
					Operation: confluence.SpacePermissionOperation2{
						Key:    strings.Join(permissionParts[:len(permissionParts)-1], "_"),
						Target: permissionParts[len(permissionParts)-1],
					},
				})
			},
			permissionRemoving: func(permission, actorId string) {
				key := fmt.Sprintf("%s:%s", actorId, permission)
				requestId := manager.permissionsMapForRemoval[key]
				_ = manager.client.SpacePermissionsService().RemovePermission(manager.spaceKey, requestId)
			},
		},
		permissionActorGroup: {
			permissionAdding: func(permission, actorId string) {
				permissionParts := strings.Split(permission, "_")
				_, _ = manager.client.SpacePermissionsService().AddPermission(manager.spaceKey, confluence.AddPermissionRequest{
					Subject: confluence.SpacePermissionSubject{
						Type: confluence.PrincipalTypeGroup,
						Id:   actorId,
					},
					Operation: confluence.SpacePermissionOperation2{
						Key:    strings.Join(permissionParts[:len(permissionParts)-1], "_"),
						Target: permissionParts[len(permissionParts)-1],
					},
				})
			},
			permissionRemoving: func(permission, actorId string) {
				key := fmt.Sprintf("%s:%s", actorId, permission)
				requestId := manager.permissionsMapForRemoval[key]
				_ = manager.client.SpacePermissionsService().RemovePermission(manager.spaceKey, requestId)
			},
		},
	}

	slices.SortFunc(permissions, confluence.SortOperation)
	for _, permission := range permissions {
		//changeReq, exists := manager.changeRequests[permission]
		//if !exists {
		//	changeReq = newPermissionChangeRequest()
		//	manager.changeRequests[permission] = changeReq
		//}

		if action := actions[permissionActor][changeType]; action != nil {
			action(permission, actorId)
		}
	}
}

//func (manager *SpacePermissionsManager) Finalized() {
//	for permission, changeReq := range manager.changeRequests {
//		projectRoleService := manager.client.SpacePermissionsService()
//		if len(changeReq.addingUsers) > 0 || len(changeReq.addingGroups) > 0 {
//			_ = projectRoleService.AddPermission(
//				manager.spaceKey,
//				roleId,
//				changeReq.addingUsers,
//				changeReq.addingGroups,
//			)
//		}
//		if len(changeReq.removingUsers) > 0 || len(changeReq.removingGroups) > 0 {
//
//			_ = projectRoleService.RemoveProjectRole(
//				manager.projectIdOrKey,
//				roleId,
//				changeReq.removingUsers,
//				changeReq.removingGroups,
//			)
//		}
//	}
//}
