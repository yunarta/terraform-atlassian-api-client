package bamboo

import (
	"fmt"
	"github.com/yunarta/golang-quality-of-life-pack/collections"
)

// A method of the ProjectService type (which should represent our project service).
// readRolePermissions fetches the Role Permission data associated with a given project key.
// It provides a unified view of both project and plan permissions for a comprehensive understanding of permissions.
func (service *ProjectService) readRolePermissions(projectKey string) (*RolePermissionResponse, error) {
	// The error return value is declared upfront, we will use it to catch and return errors as we proceed.
	var err error

	// A PermissionsHelper is used to interact with the underlying permissions service. We pass necessary context to it, like
	// Transport for communication, and the Url which identifies where to fetch role permissions from.
	// This approach to building the state necessary for interaction is simple and straightforward.
	projectPermission, err := PermissionsHelper{
		Transport: service.transport,
		Url:       fmt.Sprintf(readProjectRolePermission, projectKey),
	}.ReadRolePermissions()

	// Check for errors. Any error that occurred while trying to get permissions aborts the operation, maintaining safety
	if err != nil {
		return nil, err
	}

	// Similar to above, we get permissions for project plans involved in the project.
	projectPlanPermission, err := PermissionsHelper{
		Transport: service.transport,
		Url:       fmt.Sprintf(readProjectPlanRolePermission, projectKey),
	}.ReadRolePermissions()

	if err != nil {
		return nil, err
	}

	// Permissions of projects and project plans are combined into one map. A map provides constant time lookup,
	// making it efficient for checking if a specific permission exists or not.
	permissions := map[string]RolePermission{}
	for _, item := range projectPermission.Results {
		permissions[item.Name] = item
	}

	for _, item := range projectPlanPermission.Results {
		// Emphasize that the map gives us efficiency gains when checking if an item exists.
		_, ok := permissions[item.Name]
		if !ok {
			permissions[item.Name] = item
		} else {
			// Add new permissions to the existing ones but ensure uniqueness with the Unique function
			permission := permissions[item.Name]
			permission.Permissions = collections.Unique(append(permission.Permissions, item.Permissions...))
			permissions[item.Name] = permission
		}
	}

	// Transform the map into a slice, so it is more convenient to process (slices are easier to iterate over and manipulate).
	var permissionsAsList []RolePermission
	for _, value := range permissions {
		permissionsAsList = append(permissionsAsList, value)
	}

	// The result of our method is a 'RolePermissionResponse', which wraps the list of permissions to be easily used by recipients.
	return &RolePermissionResponse{
		Results: permissionsAsList,
	}, nil
}

func (service *ProjectService) addRolePermissions(projectKey string, role string, permissions []string) error {
	// Add a set of permissions to a specific role for a particular project.
	// This operation only does something when the role, project, and permissions are all valid.
	var err error

	var permissionScope = make([]string, 0)
	// We set the Permission Scope according to the type of User (Logged In/Anonymous).
	// This understanding of role & permission context helps target permissions updates accurately.
	if role == "LOGGED_IN" {
		permissionScope = projectPermissions
	} else {
		permissionScope = anonymousPermissions
	}

	// Similar to its usage in the previous function, PermissionsHelper makes the actual call to add permissions.
	// It contains all necessary information for doing the operation.
	err = PermissionsHelper{
		Transport:   service.transport,
		Url:         fmt.Sprintf(updateProjectRolePermission, projectKey, role),
		Permissions: collections.Intersect(permissionScope, permissions),
	}.AddPermissions()

	if err != nil {
		return err
	}

	// Repeat the process, but this time for permissions related to project plans.
	if role == "LOGGED_IN" {
		permissionScope = projectPlanPermissions
	} else {
		permissionScope = anonymousPermissions
	}

	err = PermissionsHelper{
		Transport:   service.transport,
		Url:         fmt.Sprintf(updateProjectPlanRolePermission, projectKey, role),
		Permissions: collections.Intersect(permissionScope, permissions),
	}.AddPermissions()

	return err
}

func (service *ProjectService) removeRolePermissions(projectKey string, role string, permissions []string) error {
	// Here we have a function very similar to addRolePermissions, but it removes permissions instead.
	// The same design considerations apply, again demonstrating how a design pattern can streamline code flow.
	var err error
	var permissionScope = make([]string, 0)
	if role == "LOGGED_IN" {
		permissionScope = projectPermissions
	} else {
		permissionScope = anonymousPermissions
	}

	err = PermissionsHelper{
		Transport:   service.transport,
		Url:         fmt.Sprintf(updateProjectRolePermission, projectKey, role),
		Permissions: collections.Intersect(permissionScope, permissions),
	}.RemovePermissions()

	if err != nil {
		return err
	}

	if role == "LOGGED_IN" {
		permissionScope = projectPlanPermissions
	} else {
		permissionScope = anonymousPermissions
	}

	err = PermissionsHelper{
		Transport:   service.transport,
		Url:         fmt.Sprintf(updateProjectPlanRolePermission, projectKey, role),
		Permissions: collections.Intersect(permissionScope, permissions),
	}.RemovePermissions()

	return err
}
