package bamboo

import (
	"fmt"
	"github.com/yunarta/terraform-atlassian-api-client/util"
)

// ReadPermissions function, reading permissions of a deployment, grouping them by users, groups and roles.
func (service *DeploymentService) ReadPermissions(deploymentID int) (*ObjectPermission, error) {
	// Reading group permissions for the deployment from REST API
	groupPermissions, err := service.readGroupPermissions(deploymentID, "")
	if err != nil {
		return nil, err
	}

	// Reading user permissions for the deployment from REST API
	userPermissions, err := service.readUserPermissions(deploymentID, "")
	if err != nil {
		return nil, err
	}

	// Reading role permissions for the deployment from REST API
	rolePermissions, err := service.readRolePermissions(deploymentID)
	if err != nil {
		return nil, err
	}

	// Grouping the permissions into an ObjectPermission struct
	objectPermission := ObjectPermission{
		Groups: groupPermissions.Results,
		Users:  userPermissions.Results,
		Roles:  rolePermissions.Results,
	}
	return &objectPermission, nil
}

// readGroupPermissions function, reads group permissions of a deployment.
func (service *DeploymentService) readGroupPermissions(deploymentId int, group string) (*GroupPermissionResponse, error) {
	// Returns the permissions helper to read the group permissions
	return PermissionsHelper{
		Transport: service.transport,
		Url:       fmt.Sprintf("/rest/api/latest/permissions/deployment/%d/groups?limit=1000%s", deploymentId, util.QueryParam("name", group)),
	}.ReadGroupPermissions()
}

// addGroupPermissions function, adds group permissions for a deployment.
func (service *DeploymentService) addGroupPermissions(deploymentId int, group string, permissions []string) error {
	// Returns the permissions helper to add the group permissions
	return PermissionsHelper{
		Transport:   service.transport,
		Url:         fmt.Sprintf("/rest/api/latest/permissions/deployment/%d/groups/%s", deploymentId, group),
		Permissions: permissions,
	}.AddPermissions()
}

// removeGroupPermissions function, removes group permissions for a deployment.
func (service *DeploymentService) removeGroupPermissions(deploymentId int, group string, permissions []string) error {
	// Returns the permissions helper to remove the group permissions
	return PermissionsHelper{
		Transport:   service.transport,
		Url:         fmt.Sprintf("/rest/api/latest/permissions/deployment/%d/groups/%s", deploymentId, group),
		Permissions: permissions,
	}.RemovePermissions()
}

// readUserPermissions function, reads user permissions of a deployment.
func (service *DeploymentService) readUserPermissions(deploymentId int, user string) (*UserPermissionResponse, error) {
	// Returns the permissions helper to read the user permissions
	return PermissionsHelper{
		Transport: service.transport,
		Url:       fmt.Sprintf("/rest/api/latest/permissions/deployment/%d/users?limit=1000%s", deploymentId, util.QueryParam("name", user)),
	}.ReadUserPermissions()
}

// addUserPermissions function, adds user permissions for a deployment.
func (service *DeploymentService) addUserPermissions(deploymentId int, user string, permissions []string) error {
	// Returns the permissions helper to add the user permissions
	return PermissionsHelper{
		Transport:   service.transport,
		Url:         fmt.Sprintf("/rest/api/latest/permissions/deployment/%d/users/%s", deploymentId, user),
		Permissions: permissions,
	}.AddPermissions()
}

// removeUserPermissions function, removes user permissions for a deployment.
func (service *DeploymentService) removeUserPermissions(deploymentId int, user string, permissions []string) error {
	// Returns the permissions helper to remove the user permissions
	return PermissionsHelper{
		Transport:   service.transport,
		Url:         fmt.Sprintf("/rest/api/latest/permissions/deployment/%d/users/%s", deploymentId, user),
		Permissions: permissions,
	}.RemovePermissions()
}

// readRolePermissions function, reads role permissions of a deployment.
func (service *DeploymentService) readRolePermissions(deploymentId int) (*RolePermissionResponse, error) {
	// Returns the permissions helper to read the role permissions
	return PermissionsHelper{
		Transport: service.transport,
		Url:       fmt.Sprintf("/rest/api/latest/permissions/deployment/%d/roles", deploymentId),
	}.ReadRolePermissions()
}

// addRolePermissions function, adds role permissions for a deployment.
func (service *DeploymentService) addRolePermissions(deploymentId int, role string, permissions []string) error {
	// Returns the permissions helper to add the role permissions
	return PermissionsHelper{
		Transport:   service.transport,
		Url:         fmt.Sprintf("/rest/api/latest/permissions/deployment/%d/roles/%s", deploymentId, role),
		Permissions: permissions,
	}.AddPermissions()
}

// removeRolePermissions function, removes role permissions for a deployment.
func (service *DeploymentService) removeRolePermissions(deploymentId int, role string, permissions []string) error {
	// Returns the permissions helper to remove the role permissions
	return PermissionsHelper{
		Transport:   service.transport,
		Url:         fmt.Sprintf("/rest/api/latest/permissions/deployment/%d/roles/%s", deploymentId, role),
		Permissions: permissions,
	}.RemovePermissions()
}

// UpdateGroupPermissions function, updates group permissions of a deployment.
func (service *DeploymentService) UpdateGroupPermissions(deploymentId int, group string, newPermissions []string) error {
	// Reading group permissions
	groupPermissions, err := service.readGroupPermissions(deploymentId, group)
	if err != nil {
		return err
	}

	// Updating group permissions
	return updateItemPermission(
		groupPermissions,
		deploymentId,
		group,
		newPermissions,
		service.addGroupPermissions,
		service.removeGroupPermissions,
	)
}

// UpdateUserPermissions function, updates user permissions of a deployment.
func (service *DeploymentService) UpdateUserPermissions(deploymentId int, username string, newPermissions []string) error {
	// Reading user permissions
	userPermissions, err := service.readUserPermissions(deploymentId, username)
	if err != nil {
		return err
	}

	// Updating user permissions
	return updateItemPermission(
		userPermissions,
		deploymentId,
		username,
		newPermissions,
		service.addUserPermissions,
		service.removeUserPermissions,
	)
}

// UpdateRolePermissions function, updates role permissions of a deployment.
func (service *DeploymentService) UpdateRolePermissions(deploymentId int, role string, newPermissions []string) error {
	// Reading role permissions
	rolePermissions, err := service.readRolePermissions(deploymentId)
	if err != nil {
		return err
	}

	// Updating role permissions
	return updateItemPermission(
		rolePermissions,
		deploymentId,
		role,
		newPermissions,
		service.addRolePermissions,
		service.removeRolePermissions,
	)
}

func (service *DeploymentService) findAvailableUser(deploymentId int, user string) (*UserPermissionResponse, error) {
	// Returns the permissions helper to read the user permissions
	return PermissionsHelper{
		Transport: service.transport,
		Url:       fmt.Sprintf("/rest/api/latest/permissions/deployment/%d/available-users?limit=1000%s", deploymentId, util.QueryParam("name", user)),
	}.ReadUserPermissions()
}

func (service *DeploymentService) FindAvailableUser(deploymentId int, username string) (*UserPermission, error) {
	// Reading user permissions
	userPermissions, err := service.findAvailableUser(deploymentId, username)
	if err != nil {
		return nil, err
	}

	for _, user := range userPermissions.Results {
		if user.Name == username {
			return &user, nil
		}
	}

	userPermissions, err = service.readUserPermissions(deploymentId, username)
	if err != nil {
		return nil, err
	}

	for _, user := range userPermissions.Results {
		if user.Name == username {
			return &user, nil
		}
	}

	return nil, nil
}

func (service *DeploymentService) findAvailableGroup(deploymentId int, group string) (*GroupPermissionResponse, error) {
	// Returns the permissions helper to read the user permissions
	return PermissionsHelper{
		Transport: service.transport,
		Url:       fmt.Sprintf("/rest/api/latest/permissions/deployment/%d/available-groups?limit=1000%s", deploymentId, util.QueryParam("name", group)),
	}.ReadGroupPermissions()
}

func (service *DeploymentService) FindAvailableGroup(deploymentId int, groupName string) (*GroupPermission, error) {
	// Reading user permissions
	groupPermissions, err := service.findAvailableGroup(deploymentId, groupName)
	if err != nil {
		return nil, err
	}

	for _, group := range groupPermissions.Results {
		if group.Name == groupName {
			return &group, nil
		}
	}

	groupPermissions, err = service.readGroupPermissions(deploymentId, groupName)
	if err != nil {
		return nil, err
	}

	for _, group := range groupPermissions.Results {
		if group.Name == groupName {
			return &group, nil
		}
	}

	return nil, nil
}
