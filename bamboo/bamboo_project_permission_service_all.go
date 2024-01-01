package bamboo

// Define constants
const (
	// These are endpoint templates for Bamboo API.
	// '%s' are placeholders where different parameters will be inserted to generate full API path.

	// For Project Group Permissions
	readProjectGroupPermissions  = "/rest/api/latest/permissions/project/%s/groups?limit=1000%s"
	updateProjectGroupPermission = "/rest/api/latest/permissions/project/%s/groups/%s"
	// For Project Plan Group Permissions
	readProjectPlanGroupPermissions  = "/rest/api/latest/permissions/projectplan/%s/groups?limit=1000%s"
	updateProjectPlanGroupPermission = "/rest/api/latest/permissions/projectplan/%s/groups/%s"
	// For Project User Permissions
	readProjectUserPermission   = "/rest/api/latest/permissions/project/%s/users?limit=1000%s"
	updateProjectUserPermission = "/rest/api/latest/permissions/project/%s/users/%s"
	// For Project Plan User Permissions
	readProjectPlanUserPermission    = "/rest/api/latest/permissions/projectplan/%s/users?limit=1000%s"
	updateProjectPlanUserPermissions = "/rest/api/latest/permissions/projectplan/%s/users/%s"
	// For Project Role Permissions
	readProjectRolePermission   = "/rest/api/latest/permissions/project/%s/roles"
	updateProjectRolePermission = "/rest/api/latest/permissions/project/%s/roles/%s"

	// For Project Plan Role Permissions
	readProjectPlanRolePermission   = "/rest/api/latest/permissions/projectplan/%s/roles"
	updateProjectPlanRolePermission = "/rest/api/latest/permissions/projectplan/%s/roles/%s"
)

// Define variables
var (
	// This is the list of permissions available for anonymous users
	anonymousPermissions = []string{
		"READ",
	}

	// This is the list of permissions available for a project
	projectPermissions = []string{
		"READ",
		"CREATE",
		"CREATEREPOSITORY",
		"ADMINISTRATION",
	}

	// This is the list of permissions available for a project's plan
	projectPlanPermissions = []string{
		"CLONE",
		"WRITE",
		"READ",
		"ADMINISTRATION",
		"BUILD",
		"VIEWCONFIGURATION",
	}
)

// ReadPermissions reads the permissions for a project.
func (service *ProjectService) ReadPermissions(projectKey string) (*ObjectPermission, error) {

	// Fetch Group Permissions
	groupPermissions, err := service.readGroupPermissions(projectKey, "")
	if err != nil {
		return nil, err
	}

	// Fetch User Permissions
	userPermissions, err := service.readUserPermissions(projectKey, "")
	if err != nil {
		return nil, err
	}

	// Consolidate and return Group and User Permissions
	return &ObjectPermission{
		Groups: groupPermissions.Results,
		Users:  userPermissions.Results,
	}, nil
}

// UpdateGroupPermissions updates the group permissions.
func (service *ProjectService) UpdateGroupPermissions(projectKey string, groupName string, newPermissions []string) error {
	// Fetch Existing Group Permissions
	groupPermissions, err := service.readGroupPermissions(projectKey, groupName)
	if err != nil {
		return err
	}

	// Update Permissions and return
	return updateItemPermission(
		groupPermissions,
		projectKey,
		groupName,
		newPermissions,
		service.addGroupPermissions,
		service.removeGroupPermissions)
}

// UpdateUserPermissions updates the permissions of a user.
func (service *ProjectService) UpdateUserPermissions(projectKey string, username string, newPermissions []string) error {
	// Fetch existing user permissions
	userPermissions, err := service.readUserPermissions(projectKey, username)
	if err != nil {
		return err
	}

	// Update permissions and return
	return updateItemPermission(
		userPermissions,
		projectKey,
		username,
		newPermissions,
		service.addUserPermissions,
		service.removeUserPermissions)
}

// UpdateRolePermissions updates the permissions of a role.
func (service *ProjectService) UpdateRolePermissions(projectKey string, role string, newPermissions []string) error {
	// Fetch existing role permissions
	rolePermissions, err := service.readRolePermissions(projectKey)
	if err != nil {
		return err
	}

	// Update permissions and return
	return updateItemPermission(
		rolePermissions,
		projectKey,
		role,
		newPermissions,
		service.addRolePermissions,
		service.removeRolePermissions)
}
