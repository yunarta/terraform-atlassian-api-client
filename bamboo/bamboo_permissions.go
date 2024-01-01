package bamboo

// PermissionResponse interface defines a method for finding a PermissionsComparator based on a key.
type PermissionResponse interface {
	Find(key string) PermissionsComparator
}

// PermissionsComparator interface defines a method for comparing permissions.
type PermissionsComparator interface {
	DeltaPermissions(permissions []string) ([]string, []string)
}

// ObjectPermission struct holds permission data for groups, users, and roles.
type ObjectPermission struct {
	Groups []GroupPermission
	Users  []UserPermission
	Roles  []RolePermission
}

// GroupPermissionResponse struct represents the response format for group permissions.
type GroupPermissionResponse struct {
	Start   int               `json:"start,omitempty"`   // Starting index of the result set.
	Limit   int               `json:"limit,omitempty"`   // Limit on the number of results.
	Results []GroupPermission `json:"results,omitempty"` // Actual group permissions.
}

// Ensure GroupPermissionResponse implements PermissionResponse.
var _ PermissionResponse = &GroupPermissionResponse{}

// Find locates a group by name and returns it as a PermissionsComparator.
func (permissions *GroupPermissionResponse) Find(groupName string) PermissionsComparator {
	for _, permission := range permissions.Results {
		if permission.Name == groupName {
			return &permission
		}
	}
	return nil
}

// GroupPermission struct represents a group's permissions.
type GroupPermission struct {
	Name        string   `json:"name,omitempty"`        // Name of the group.
	Editable    bool     `json:"editable,omitempty"`    // Indicates if the permission is editable.
	Permissions []string `json:"permissions,omitempty"` // List of permissions.
}

// Ensure GroupPermission implements PermissionsComparator.
var _ PermissionsComparator = &GroupPermission{}

// DeltaPermissions compares current permissions with new ones, returning additions and removals.
func (g *GroupPermission) DeltaPermissions(newPermissions []string) ([]string, []string) {
	return deltaPermissions(g.Permissions, newPermissions)
}

// UserPermissionResponse struct represents the response format for user permissions.
type UserPermissionResponse struct {
	Start   int              `json:"start,omitempty"`
	Limit   int              `json:"limit,omitempty"`
	Results []UserPermission `json:"results,omitempty"`
}

// Ensure UserPermissionResponse implements PermissionResponse.
var _ PermissionResponse = &UserPermissionResponse{}

// Find locates a user by name and returns it as a PermissionsComparator.
func (permissions *UserPermissionResponse) Find(userName string) PermissionsComparator {
	for _, permission := range permissions.Results {
		if permission.Name == userName {
			return &permission
		}
	}
	return nil
}

// UserPermission struct represents a user's permissions.
type UserPermission struct {
	Name        string   `json:"name,omitempty"`
	FullName    string   `json:"fullName,omitempty"`
	Email       string   `json:"email,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}

// Ensure UserPermission implements PermissionsComparator.
var _ PermissionsComparator = &UserPermission{}

// DeltaPermissions compares current permissions with new ones, returning additions and removals.
func (u *UserPermission) DeltaPermissions(newPermissions []string) ([]string, []string) {
	return deltaPermissions(u.Permissions, newPermissions)
}

// RolePermissionResponse struct represents the response format for role permissions.
type RolePermissionResponse struct {
	Start   int              `json:"start,omitempty"`
	Limit   int              `json:"limit,omitempty"`
	Results []RolePermission `json:"results,omitempty"`
}

// Ensure RolePermissionResponse implements PermissionResponse.
var _ PermissionResponse = &RolePermissionResponse{}

// Find locates a role by name and returns it as a PermissionsComparator.
func (permissions *RolePermissionResponse) Find(roleName string) PermissionsComparator {
	for _, permission := range permissions.Results {
		if permission.Name == roleName {
			return &permission
		}
	}

	return nil
}

// RolePermission struct represents a role's permissions.
type RolePermission struct {
	Name        string   `json:"name,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}

// Ensure RolePermission implements PermissionsComparator.
var _ PermissionsComparator = &RolePermission{}

// DeltaPermissions compares current permissions with new ones, returning additions and removals.
func (u *RolePermission) DeltaPermissions(newPermissions []string) ([]string, []string) {
	return deltaPermissions(u.Permissions, newPermissions)
}

// deltaPermissions identifies which permissions are being added and which are being removed.
func deltaPermissions(existingPermissions []string, newPermissions []string) ([]string, []string) {
	var (
		addingPermissions   = make([]string, 0) // Permissions to add.
		removingPermissions = make([]string, 0) // Permissions to remove.
	)

	// Loop to find new permissions not in existing permissions.
	for _, newPermission := range newPermissions {
		if !contains(existingPermissions, newPermission) {
			addingPermissions = append(addingPermissions, newPermission)
		}
	}

	// Loop to find existing permissions not in new permissions.
	for _, existingPermission := range existingPermissions {
		if !contains(newPermissions, existingPermission) {
			removingPermissions = append(removingPermissions, existingPermission)
		}
	}

	return addingPermissions, removingPermissions
}

// contains checks if a permission exists in a slice of permissions.
func contains(slice []string, perm string) bool {
	for _, v := range slice {
		if v == perm {
			return true
		}
	}
	return false
}
