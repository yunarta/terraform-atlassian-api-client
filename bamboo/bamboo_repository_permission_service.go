package bamboo

import (
	"fmt"
	"github.com/yunarta/terraform-atlassian-api-client/util"
)

// ReadPermissions retrieves permissions for a specific repository.
// It combines group, user, and role permissions into a single ObjectPermission.
func (service *RepositoryService) ReadPermissions(repositoryId int) (*ObjectPermission, error) {
	// Read group permissions
	groupPermissions, err := service.readGroupPermissions(repositoryId, "")
	if err != nil {
		return nil, err
	}

	// Read user permissions
	userPermissions, err := service.readUserPermissions(repositoryId, "")
	if err != nil {
		return nil, err
	}

	// Read role permissions
	rolePermissions, err := service.readRolePermissions(repositoryId)
	if err != nil {
		return nil, err
	}

	// Aggregate permissions into one object
	objectPermission := ObjectPermission{
		Groups: groupPermissions.Results,
		Users:  userPermissions.Results,
		Roles:  rolePermissions.Results,
	}

	return &objectPermission, nil
}

// readGroupPermissions is a helper function to retrieve group permissions from a repository.
func (service *RepositoryService) readGroupPermissions(repositoryId int, group string) (*GroupPermissionResponse, error) {
	// Constructs and uses a PermissionsHelper to read group permissions
	return PermissionsHelper{
		Transport: service.transport,
		Url:       fmt.Sprintf("/rest/api/latest/permissions/repository/%d/groups?limit=1000%s", repositoryId, util.QueryParam("name", group)),
	}.ReadGroupPermissions()
}

// addGroupPermissions is a helper function to add permissions for a group in a repository.
func (service *RepositoryService) addGroupPermissions(repositoryId int, group string, permissions []string) error {
	// Constructs and uses a PermissionsHelper to add group permissions
	return PermissionsHelper{
		Transport:   service.transport,
		Url:         fmt.Sprintf("/rest/api/latest/permissions/repository/%d/groups/%s", repositoryId, group),
		Permissions: permissions,
	}.AddPermissions()
}

// removeGroupPermissions is a helper function to remove permissions for a group in a repository.
func (service *RepositoryService) removeGroupPermissions(repositoryId int, group string, permissions []string) error {
	// Constructs and uses a PermissionsHelper to remove group permissions
	return PermissionsHelper{
		Transport:   service.transport,
		Url:         fmt.Sprintf("/rest/api/latest/permissions/repository/%d/groups/%s", repositoryId, group),
		Permissions: permissions,
	}.RemovePermissions()
}

// readUserPermissions is a helper function to retrieve user permissions from a repository.
func (service *RepositoryService) readUserPermissions(repositoryId int, user string) (*UserPermissionResponse, error) {
	// Constructs and uses a PermissionsHelper to read user permissions
	return PermissionsHelper{
		Transport: service.transport,
		Url:       fmt.Sprintf("/rest/api/latest/permissions/repository/%d/users?limit=1000%s", repositoryId, util.QueryParam("name", user)),
	}.ReadUserPermissions()
}

// addUserPermissions is a helper function to add permissions for a user in a repository.
func (service *RepositoryService) addUserPermissions(repositoryId int, user string, permissions []string) error {
	// Constructs and uses a PermissionsHelper to add user permissions
	return PermissionsHelper{
		Transport:   service.transport,
		Url:         fmt.Sprintf("/rest/api/latest/permissions/repository/%d/users/%s", repositoryId, user),
		Permissions: permissions,
	}.AddPermissions()
}

// removeUserPermissions is a helper function to remove permissions for a user in a repository.
func (service *RepositoryService) removeUserPermissions(repositoryId int, user string, permissions []string) error {
	// Constructs and uses a PermissionsHelper to remove user permissions
	return PermissionsHelper{
		Transport:   service.transport,
		Url:         fmt.Sprintf("/rest/api/latest/permissions/repository/%d/users/%s", repositoryId, user),
		Permissions: permissions,
	}.RemovePermissions()
}

// readRolePermissions is a helper function to retrieve role permissions from a repository.
func (service *RepositoryService) readRolePermissions(repositoryId int) (*RolePermissionResponse, error) {
	// Constructs and uses a PermissionsHelper to read role permissions
	return PermissionsHelper{
		Transport: service.transport,
		Url:       fmt.Sprintf("/rest/api/latest/permissions/repository/%d/roles", repositoryId),
	}.ReadRolePermissions()
}

// addRolePermissions is a helper function to add permissions for a role in a repository.
func (service *RepositoryService) addRolePermissions(repositoryId int, role string, permissions []string) error {
	// Constructs and uses a PermissionsHelper to add role permissions
	return PermissionsHelper{
		Transport:   service.transport,
		Url:         fmt.Sprintf("/rest/api/latest/permissions/repository/%d/roles/%s", repositoryId, role),
		Permissions: permissions,
	}.AddPermissions()
}

// removeRolePermissions is a helper function to remove permissions for a role in a repository.
func (service *RepositoryService) removeRolePermissions(repositoryId int, role string, permissions []string) error {
	// Constructs and uses a PermissionsHelper to remove role permissions
	return PermissionsHelper{
		Transport:   service.transport,
		Url:         fmt.Sprintf("/rest/api/latest/permissions/repository/%d/roles/%s", repositoryId, role),
		Permissions: permissions,
	}.RemovePermissions()
}

// UpdateGroupPermissions updates the permissions of a group within a repository.
// It handles the addition and removal of permissions based on the current state and the desired new permissions.
func (service *RepositoryService) UpdateGroupPermissions(repositoryId int, group string, newPermissions []string) error {
	groupPermissions, err := service.readGroupPermissions(repositoryId, group)
	if err != nil {
		return err
	}

	// Calls a generic function to update permissions
	return updateItemPermission(
		groupPermissions,
		repositoryId,
		group,
		newPermissions,
		service.addGroupPermissions,
		service.removeGroupPermissions,
	)
}

// UpdateUserPermissions updates the permissions of a user within a repository.
// Similar to UpdateGroupPermissions, it handles the changes based on current and new permissions.
func (service *RepositoryService) UpdateUserPermissions(repositoryId int, username string, newPermissions []string) error {
	userPermissions, err := service.readUserPermissions(repositoryId, username)
	if err != nil {
		return err
	}

	// Calls a generic function to update permissions
	return updateItemPermission(
		userPermissions,
		repositoryId,
		username,
		newPermissions,
		service.addUserPermissions,
		service.removeUserPermissions,
	)
}

// UpdateRolePermissions updates the permissions of a role within a repository.
// It follows the same pattern as updating group and user permissions.
func (service *RepositoryService) UpdateRolePermissions(repositoryId int, role string, newPermissions []string) error {
	rolePermissions, err := service.readRolePermissions(repositoryId)
	if err != nil {
		return err
	}

	// Calls a generic function to update permissions
	return updateItemPermission(
		rolePermissions,
		repositoryId,
		role,
		newPermissions,
		service.addRolePermissions,
		service.removeRolePermissions,
	)
}
