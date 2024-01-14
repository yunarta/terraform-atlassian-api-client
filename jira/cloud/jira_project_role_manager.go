package cloud

import (
	"fmt"
	"github.com/yunarta/golang-quality-of-life-pack/collections"
	"github.com/yunarta/terraform-atlassian-api-client/jira"
)

type RoleActorType int
type RoleChangeType int

const (
	roleActorUser RoleActorType = iota
	roleActorGroup
)

const (
	roleAdding RoleChangeType = iota
	roleRemoving
)

// ProjectRoleManage manages project roles within a Jira project.
type ProjectRoleManage struct {
	client         *JiraClient
	projectIdOrKey string
	ReadOnly       bool

	objectRoles *jira.ObjectRoles

	lookupRole     map[string]string             // map[role.name]role.id
	changeRequests map[string]*roleChangeRequest // map[role.id]changeRequest
}

// roleChangeRequest stores changes to be made to a user's or group's roles.
type roleChangeRequest struct {
	addingUsers    []string
	removingUsers  []string
	addingGroups   []string
	removingGroups []string
}

// newRoleChangeRequest initializes a new roleChangeRequest.
func newRoleChangeRequest() *roleChangeRequest {
	return &roleChangeRequest{
		addingUsers:    make([]string, 0),
		removingUsers:  make([]string, 0),
		addingGroups:   make([]string, 0),
		removingGroups: make([]string, 0),
	}
}

// NewProjectRoleManager creates a new instance of ProjectRoleManage.
func NewProjectRoleManager(client *JiraClient, projectIdOrKey string) *ProjectRoleManage {
	return &ProjectRoleManage{
		client:         client,
		projectIdOrKey: projectIdOrKey,
		lookupRole:     make(map[string]string),
		changeRequests: make(map[string]*roleChangeRequest),
	}
}

// ReadRoles retrieves and processes roles for the project.
func (manager *ProjectRoleManage) ReadRoles(allRoles []string) (*jira.ObjectRoles, error) {
	projectRoles, err := manager.client.ProjectRoleService().ReadProjectRoles(manager.projectIdOrKey)
	if err != nil {
		return nil, err
	}

	filteredRoles := manager.filterRolesByNames(projectRoles, allRoles)
	groupRoles, userRoles := manager.processRoles(filteredRoles)

	accountIds, userRolesList := collections.SpliceMapToKeyValue(userRoles)
	groupIds, groupRolesList := collections.SpliceMapToKeyValue(groupRoles)

	if !manager.ReadOnly {
		// We register the account ids and group ids for caching purpose
		manager.client.ActorLookupService().RegisterAccountIds(accountIds...)
		manager.client.ActorLookupService().RegisterGroupIds(groupIds...)
	}

	manager.objectRoles = &jira.ObjectRoles{
		Groups: groupRolesList,
		Users:  userRolesList,
	}
	return manager.objectRoles, nil
}

// filterRolesByNames filters roles based on provided names.
func (manager *ProjectRoleManage) filterRolesByNames(roles []jira.RoleType, roleNames []string) []jira.RoleType {
	// Then we only focus on roles that in stated in the parameters
	// This is for REST API optimization
	filteredRoles := make([]jira.RoleType, 0)
	for _, role := range roles {
		if collections.Contains(roleNames, role.Name) {
			filteredRoles = append(filteredRoles, role)
			manager.lookupRole[role.Name] = role.ID
		}
	}
	return filteredRoles
}

// processRoles organizes role data into user and group mappings.
func (manager *ProjectRoleManage) processRoles(roles []jira.RoleType) (map[string]jira.GroupRoles, map[string]jira.UserRoles) {
	groupRoles := make(map[string]jira.GroupRoles)
	userRoles := make(map[string]jira.UserRoles)

	for _, role := range roles {
		actors, err := manager.client.ProjectRoleService().ReadProjectRoleActors(manager.projectIdOrKey, role.ID)
		if err != nil {
			continue // Log the error in production code
		}
		for _, actor := range actors {
			if actor.Type == "atlassian-group-role-actor" {
				manager.addGroupRole(groupRoles, role.Name, actor)
			} else if actor.Type == "atlassian-user-role-actor" {
				manager.addUserRole(userRoles, role.Name, actor)
			}
		}
	}
	return groupRoles, userRoles
}

// addGroupRole adds a role to a group.
func (manager *ProjectRoleManage) addGroupRole(groupRoles map[string]jira.GroupRoles, roleName string, actor jira.Actor) {
	// First we retrieve the group from the map if available
	groupId := actor.ActorGroup.GroupId
	group, exists := groupRoles[groupId]
	if !exists {
		// The group not in the map, so we add the group for the first time
		group = jira.GroupRoles{
			Name:      actor.DisplayName,
			AccountId: groupId,
			Roles:     []string{},
		}
	}
	group.Roles = append(group.Roles, roleName)
	groupRoles[groupId] = group
}

// addUserRole adds a role to a user.
func (manager *ProjectRoleManage) addUserRole(userRoles map[string]jira.UserRoles, roleName string, actor jira.Actor) {
	// First we retrieve the user from the map if available

	accountId := actor.ActorUser.AccountID
	user, exists := userRoles[accountId]
	if !exists {
		name := actor.DisplayName
		jiraUser := manager.client.ActorLookupService().FindUserById(accountId)
		if jiraUser != nil {
			// found system user
			name = jiraUser.EmailAddress
		}

		user = jira.UserRoles{
			Name:      name,
			AccountId: accountId,
			Roles:     []string{},
		}
	}
	user.Roles = append(user.Roles, roleName)
	userRoles[accountId] = user
}

// UpdateUserRoles updates roles assigned to a user.
func (manager *ProjectRoleManage) UpdateUserRoles(username string, newRoles []string) error {
	// read assigned roles for selected group
	jiraUser := manager.client.ActorLookupService().FindUser(username)
	if jiraUser == nil {
		return fmt.Errorf("unable to find user %s", username)
	}

	user := manager.objectRoles.FindUser(jiraUser.AccountID)

	if user != nil {
		adding, removing := user.DeltaRoles(newRoles)
		manager.prepareRoleChanges(roleActorUser, jiraUser.AccountID, adding, removing)
	} else if len(newRoles) > 0 {
		// If the item is not found but there are new permissions, add them.
		manager.makeRoleChange(roleActorUser, jiraUser.AccountID, roleAdding, newRoles)
	}
	return nil
}

// UpdateGroupRoles updates roles assigned to a group.
func (manager *ProjectRoleManage) UpdateGroupRoles(groupName string, newRoles []string) error {
	// read assigned roles for selected group
	jiraGroup := manager.client.ActorLookupService().FindGroup(groupName)
	if jiraGroup == nil {
		return fmt.Errorf("unable to find group %s", jiraGroup)
	}

	group := manager.objectRoles.FindGroup(jiraGroup.GroupId)

	if group != nil {
		adding, removing := group.DeltaRoles(newRoles)
		manager.prepareRoleChanges(roleActorGroup, jiraGroup.GroupId, adding, removing)
	} else if len(newRoles) > 0 {
		// If the item is not found but there are new permissions, add them.
		manager.makeRoleChange(roleActorGroup, jiraGroup.GroupId, roleAdding, newRoles)
	}
	return nil
}

// prepareRoleChanges prepares changes to be made to roles.
func (manager *ProjectRoleManage) prepareRoleChanges(roleActor RoleActorType, actorId string, adding, removing []string) {
	if len(adding) > 0 {
		// Add new permissions if there are any to add.
		manager.makeRoleChange(roleActor, actorId, roleAdding, adding)
	}
	if len(removing) > 0 {
		// Remove old permissions if there are any to remove.
		manager.makeRoleChange(roleActor, actorId, roleRemoving, removing)
	}
}

// makeRoleChange records a pending change to a role.
func (manager *ProjectRoleManage) makeRoleChange(roleActor RoleActorType, actorId string, changeType RoleChangeType, roles []string) {
	actions := map[RoleActorType]map[RoleChangeType]func(*roleChangeRequest, string){
		roleActorUser: {
			roleAdding: func(req *roleChangeRequest, actorId string) {
				req.addingUsers = append(req.addingUsers, actorId)
			},
			roleRemoving: func(req *roleChangeRequest, actorId string) {
				req.removingUsers = append(req.removingUsers, actorId)
			},
		},
		roleActorGroup: {
			roleAdding: func(req *roleChangeRequest, actorId string) {
				req.addingGroups = append(req.addingGroups, actorId)
			},
			roleRemoving: func(req *roleChangeRequest, actorId string) {
				req.removingGroups = append(req.removingGroups, actorId)
			},
		},
	}

	for _, role := range roles {
		roleId := manager.lookupRole[role]
		changeReq, exists := manager.changeRequests[roleId]
		if !exists {
			changeReq = newRoleChangeRequest()
			manager.changeRequests[roleId] = changeReq
		}

		if action := actions[roleActor][changeType]; action != nil {
			action(changeReq, actorId)
		}
	}
}

// Finalized applies all pending role changes to the project.
func (manager *ProjectRoleManage) Finalized() {
	for roleId, changeReq := range manager.changeRequests {
		projectRoleService := manager.client.ProjectRoleService()
		if len(changeReq.addingUsers) > 0 || len(changeReq.addingGroups) > 0 {
			_ = projectRoleService.AddProjectRole(
				manager.projectIdOrKey,
				roleId,
				changeReq.addingUsers,
				changeReq.addingGroups,
			)
		}

		if len(changeReq.removingUsers) > 0 || len(changeReq.removingGroups) > 0 {
			_ = projectRoleService.RemoveProjectRole(
				manager.projectIdOrKey,
				roleId,
				changeReq.removingUsers,
				changeReq.removingGroups,
			)
		}
	}
}
