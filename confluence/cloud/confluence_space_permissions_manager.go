package cloud

import (
	"fmt"
	"github.com/yunarta/golang-quality-of-life-pack/collections"
	"github.com/yunarta/terraform-atlassian-api-client/confluence"
	"slices"
	"strings"
)

type permissionActorType int
type permissionsChangeType int

const (
	permissionActorUser permissionActorType = iota
	permissionActorGroup
)

const (
	permissionAdding permissionsChangeType = iota
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

func NewSpaceRoleManager(client *ConfluenceClient, spaceKey string) *SpacePermissionsManager {
	return &SpacePermissionsManager{
		spaceKey:       spaceKey,
		client:         client,
		changeRequests: make(map[string]*permissionChangeRequest),
	}
}

func (manager *SpacePermissionsManager) ReadPermissions() (*confluence.ObjectPermissions, error) {
	space, err := manager.client.SpaceService().Read(manager.spaceKey)
	if err != nil {
		return nil, err
	}

	permissions, err := manager.client.SpacePermissionsService().Read(space.Id)
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
		key := fmt.Sprintf("%s:%s", permission.Principal.Id, permission.Operation.GetSlug())
		permissionsMapForRemoval[key] = permission.Id
		if permission.Principal.Type == confluence.PrincipalUser {
			username := manager.client.ActorLookupService().FindUserById(permission.Principal.Id)
			manager.addUserRole(userRoles, permission.Operation, permission.Principal.Id, username)

		} else if permission.Principal.Type == confluence.PrincipalGroup {
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

func (manager *SpacePermissionsManager) addGroupRole(groupRoles map[string]confluence.GroupPermissions, permission confluence.OperationV2, principalId, name string) {
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
	group.Permissions = append(group.Permissions, permission.GetSlug())
	groupRoles[principalId] = group
}

func (manager *SpacePermissionsManager) addUserRole(userRoles map[string]confluence.UserPermissions, permission confluence.OperationV2, principalId, name string) {
	// First we retrieve the user from the map if available
	user, exists := userRoles[principalId]
	if !exists {
		user = confluence.UserPermissions{
			Name:        name,
			AccountId:   principalId,
			Permissions: []string{},
		}
	}
	user.Permissions = append(user.Permissions, permission.GetSlug())
	userRoles[principalId] = user
}

func (manager *SpacePermissionsManager) getUserGroupIds(permissions *[]confluence.PermissionV2) ([]string, []string) {
	usersMap := make(map[string]bool)
	groupsMap := make(map[string]bool)

	for _, permission := range *permissions {
		if permission.Principal.Type == confluence.PrincipalUser {
			usersMap[permission.Principal.Id] = true
		} else if permission.Principal.Type == confluence.PrincipalGroup {
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
		manager.preparePermissionChanges(permissionActorUser, accountId, adding, removing)
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
		manager.preparePermissionChanges(permissionActorGroup, groupId, adding, removing)
	} else if len(newRoles) > 0 {
		// If the item is not found but there are new permissions, add them.
		manager.makePermissionChange(permissionActorGroup, groupId, permissionAdding, newRoles)
	}
	return nil
}

func (manager *SpacePermissionsManager) preparePermissionChanges(permissionActor permissionActorType, actorId string, adding, removing []string) {
	if len(adding) > 0 {
		// Add new permissions if there are any to add.
		manager.makePermissionChange(permissionActor, actorId, permissionAdding, adding)
	}
	if len(removing) > 0 {
		// Remove old permissions if there are any to remove.
		manager.makePermissionChange(permissionActor, actorId, permissionRemoving, removing)
	}
}

func (manager *SpacePermissionsManager) makePermissionChange(permissionActor permissionActorType, actorId string, changeType permissionsChangeType, permissions []string) {
	actions := map[permissionActorType]map[permissionsChangeType]func(string, string){
		permissionActorUser: {
			permissionAdding:   manager.addPermission,
			permissionRemoving: manager.removePermission,
		},
		permissionActorGroup: {
			permissionAdding:   manager.addPermission,
			permissionRemoving: manager.removePermission,
		},
	}

	slices.SortFunc(permissions, confluence.SortOperation)
	for _, permission := range permissions {
		if action := actions[permissionActor][changeType]; action != nil {
			action(permission, actorId)
		}
	}
}

func (manager *SpacePermissionsManager) addPermission(permission, actorId string) {
	permissionParts := strings.Split(permission, "_")
	_, _ = manager.client.SpacePermissionsService().Create(manager.spaceKey, confluence.AddPermission{
		Subject: confluence.Subject{
			Type: confluence.PrincipalUser,
			Id:   actorId,
		},
		Operation: confluence.AddOperation{
			Key:    strings.Join(permissionParts[:len(permissionParts)-1], "_"),
			Target: permissionParts[len(permissionParts)-1],
		},
	})
}

func (manager *SpacePermissionsManager) removePermission(permission, actorId string) {
	key := fmt.Sprintf("%s:%s", actorId, permission)
	requestId := manager.permissionsMapForRemoval[key]
	_ = manager.client.SpacePermissionsService().Delete(manager.spaceKey, requestId)
}
