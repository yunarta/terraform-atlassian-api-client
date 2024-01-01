package bamboo

import (
	"fmt"
	"github.com/yunarta/golang-quality-of-life-pack/collections"
	"github.com/yunarta/terraform-atlassian-api-client/util"
)

// readUserPermissions reads all project-level and plan-level permissions associated with a given user in a specific project,
// removing any duplicates and formatting them into an easily manageable list. This function is crucial
// in implementing a user permissions system.
func (service *ProjectService) readUserPermissions(projectKey string, user string) (*UserPermissionResponse, error) {

	var err error // Stores any error that we may encounter during the process.

	// We get project-level permissions for the user and store them in 'projectPermission'
	// The PermissionsHelper struct is used to perform API calls. The helper requires:
	// - Transport: A method of sending the API request.
	// - Url: The endpoint of our API request. Here it is generated using Sprintf.
	projectPermission, err := PermissionsHelper{
		Transport: service.transport,
		Url:       fmt.Sprintf(readProjectUserPermission, projectKey, util.QueryParam("name", user)),
	}.ReadUserPermissions()

	// If we get an error at this stage, we immediately return it to the caller.
	if err != nil {
		return nil, err
	}

	// We do the same thing as above but now for plan-level permissions for the user.
	projectPlanPermission, err := PermissionsHelper{
		Transport: service.transport,
		Url:       fmt.Sprintf(readProjectPlanUserPermission, projectKey, util.QueryParam("name", user)),
	}.ReadUserPermissions()

	if err != nil {
		return nil, err
	}

	// Here we create an empty map (like a dictionary in Python) to hold all the permissions.
	// The key of the map will be the permission name, and the value will be the permission details.
	permissions := map[string]UserPermission{}
	for _, item := range projectPermission.Results {
		permissions[item.Name] = item
	}

	// Next, we add the plan-level permissions to our permissions map.
	// Notice that we have a safety check to exclude any permissions that are already present in the map.
	for _, item := range projectPlanPermission.Results {
		if _, ok := permissions[item.Name]; !ok {
			permissions[item.Name] = item
		}

		// Here we ensure that for each permission name, the list of permissions is unique. No duplicates.
		permission := permissions[item.Name]
		permission.Permissions = collections.Unique(append(permission.Permissions, item.Permissions...))
		permissions[item.Name] = permission
	}

	// Finally, we transform our map into a list of permissions and return it to the caller.
	// Lists are easier to traverse and manipulate compared to maps, especially for beginners.
	var permissionsAsList []UserPermission
	for _, value := range permissions {
		permissionsAsList = append(permissionsAsList, value)
	}

	return &UserPermissionResponse{
		Results: permissionsAsList,
	}, nil
}

// addUserPermissions function tries to add user permissions specific to a project and its plan.
// Remember that permissions usually enable a user to do certain things, such as deleting or editing a project.
func (service *ProjectService) addUserPermissions(projectKey string, username string, permissions []string) error {
	var err error

	// First, we add the project specific permissions using PermissionHelper.
	// The permissions added are the intersection of existing 'projectPermissions' and new 'permissions'.
	// This is to ensure only valid permissions are added.
	err = PermissionsHelper{
		Transport:   service.transport,
		Url:         fmt.Sprintf(updateProjectUserPermission, projectKey, username),
		Permissions: collections.Intersect(projectPermissions, permissions),
	}.AddPermissions()

	// If we encounter an error, we immediately return it to the caller.
	if err != nil {
		return err
	}

	// Next, we add the plan specific permissions, again using PermissionsHelper.
	// The process is similar to adding the project specific permissions.
	err = PermissionsHelper{
		Transport:   service.transport,
		Url:         fmt.Sprintf(updateProjectPlanUserPermissions, projectKey, username),
		Permissions: collections.Intersect(projectPlanPermissions, permissions),
	}.AddPermissions()

	// Finally, return any error we might have encountered (nil in case of no error)
	return err
}

// removeUserPermissions operates similarly to 'addUserPermissions' but instead of adding permissions,
// it removes them. It's an important function to have in our system for revoking accesses.
func (service *ProjectService) removeUserPermissions(projectKey string, username string, permissions []string) error {
	var err error

	// Removes project specific permissions from the user. The process is similar to adding permissions above.
	err = PermissionsHelper{
		Transport:   service.transport,
		Url:         fmt.Sprintf(updateProjectUserPermission, projectKey, username),
		Permissions: collections.Intersect(projectPermissions, permissions),
	}.RemovePermissions()

	if err != nil {
		return err
	}

	// Removes project plan specific permissions from the user. Again, the process is similar to adding permissions.
	err = PermissionsHelper{
		Transport:   service.transport,
		Url:         fmt.Sprintf(updateProjectPlanUserPermissions, projectKey, username),
		Permissions: collections.Intersect(projectPlanPermissions, permissions),
	}.RemovePermissions()

	// Return any error we might have encountered (nil in case of no error).
	return err
}
