package bitbucket

type PermissionOwner struct {
	Name string `json:"name,omitempty"`
}

type ObjectPermission struct {
	Groups []GroupPermission
	Users  []UserPermission
}

type GroupPermissionResponse struct {
	Start  int               `json:"start,omitempty"`
	Limit  int               `json:"limit,omitempty"`
	Values []GroupPermission `json:"values,omitempty"`
}

func (permissions *GroupPermissionResponse) Find(groupName string) *GroupPermission {
	for _, permission := range permissions.Values {
		if permission.Owner.Name == groupName {
			return &permission
		}
	}
	return nil
}

type GroupPermission struct {
	Owner      PermissionOwner `json:"group,omitempty"`
	Permission string          `json:"permission,omitempty"`
}

type UserPermissionResponse struct {
	Start  int              `json:"start,omitempty"`
	Limit  int              `json:"limit,omitempty"`
	Values []UserPermission `json:"values,omitempty"`
}

func (permissions *UserPermissionResponse) Find(user string) *UserPermission {
	for _, permission := range permissions.Values {
		if permission.Owner.Name == user {
			return &permission
		}
	}
	return nil
}

type UserPermission struct {
	Owner      PermissionOwner `json:"user,omitempty"`
	Permission string          `json:"permission,omitempty"`
}
