package bamboo

import (
	"fmt"
	"github.com/yunarta/golang-quality-of-life-pack/collections"
	"github.com/yunarta/terraform-atlassian-api-client/util"
)

// ProjectService is a struct with methods to help in managing project permissions

// readGroupPermissions function reads a group's permissions based on the project key and group name
// projectKey and groupName are the input parameters, and it returns GroupPermissionResponse and error if any
// GroupPermissionResponse is a struct with the relevant information about the permissions of a group
func (service *ProjectService) readGroupPermissions(projectKey string, groupName string) (*GroupPermissionResponse, error) {
	// Initializing an error
	var err error

	// Here, call to the helper function ReadGroupPermissions is made. If an error is encountered, it is returned immediately.
	projectPermission, err := PermissionsHelper{
		Transport: service.transport,
		Url:       fmt.Sprintf(readProjectGroupPermissions, projectKey, util.QueryParam("name", groupName)),
	}.ReadGroupPermissions()

	if err != nil {
		return nil, err
	}

	// Another call to the ReadGroupPermissions for projectPlanPermission. Again, if an error is encountered, it is immediately returned.
	projectPlanPermission, err := PermissionsHelper{
		Transport: service.transport,
		Url:       fmt.Sprintf(readProjectPlanGroupPermissions, projectKey, util.QueryParam("name", groupName)),
	}.ReadGroupPermissions()

	if err != nil {
		return nil, err
	}

	// permissions is a map of GroupPermission objects. GroupPermission represent the permissions a group has.
	permissions := map[string]GroupPermission{}

	// Looping over the results from projectPermission
	for _, item := range projectPermission.Results {
		permissions[item.Name] = item
	}

	// Looping over the results from projectPlanPermission. If the permission isn't already in permissions, it's added.
	// If it's already there, the permissions are combined
	for _, item := range projectPlanPermission.Results {
		// If the permission item is not in the map, add it to the map.
		_, ok := permissions[item.Name]

		if !ok {
			permissions[item.Name] = item
		}

		// Read the permission, update it and put it back into the map.
		permission := permissions[item.Name]
		permission.Permissions = collections.Unique(append(permission.Permissions, item.Permissions...))
		permissions[item.Name] = permission
	}

	// Map is then converted to a list
	var permissionsAsList []GroupPermission
	for _, value := range permissions {
		permissionsAsList = append(permissionsAsList, value)
	}

	// Returns the list of permissions
	return &GroupPermissionResponse{
		Results: permissionsAsList,
	}, nil
}

// addGroupPermissions function adds a group's permissions based on the project key, username and list of permissions
// It returns an error if encountered during the operation
func (service *ProjectService) addGroupPermissions(projectKey string, username string, permissions []string) error {
	// Initializing an error
	var err error

	// Calling the helper function AddPermissions.
	// If an error is encountered, it's returned immediately
	err = PermissionsHelper{
		Transport:   service.transport,
		Url:         fmt.Sprintf(updateProjectGroupPermission, projectKey, username),
		Permissions: collections.Intersect(projectPermissions, permissions),
	}.AddPermissions()

	if err != nil {
		return err
	}

	// Again calling AddPermissions for project plan permissions update
	err = PermissionsHelper{
		Transport:   service.transport,
		Url:         fmt.Sprintf(updateProjectPlanGroupPermission, projectKey, username),
		Permissions: collections.Intersect(projectPlanPermissions, permissions),
	}.AddPermissions()

	// return the error, if any
	return err
}

// removeGroupPermissions function removes a group's permissions based on the project key, username and list of permissions
// It returns an error if encountered during the operation
func (service *ProjectService) removeGroupPermissions(projectKey string, username string, permissions []string) error {
	// Initializing an error
	var err error

	// Calling the helper function RemovePermissions.
	// If an error is encountered, it's returned immediately
	err = PermissionsHelper{
		Transport:   service.transport,
		Url:         fmt.Sprintf(updateProjectGroupPermission, projectKey, username),
		Permissions: collections.Intersect(projectPermissions, permissions),
	}.RemovePermissions()

	if err != nil {
		return err
	}

	// Again calling RemovePermissions for project plan permissions update
	err = PermissionsHelper{
		Transport:   service.transport,
		Url:         fmt.Sprintf(updateProjectPlanGroupPermission, projectKey, username),
		Permissions: collections.Intersect(projectPlanPermissions, permissions),
	}.RemovePermissions()

	// return the error, if any
	return err
}
