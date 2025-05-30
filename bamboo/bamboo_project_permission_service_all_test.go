package bamboo

import (
	"github.com/stretchr/testify/assert"
	"github.com/yunarta/terraform-api-transport/transport"
	"slices"
	"strings"
	"testing"
)

func TestProjectService_ReadPermissions(t *testing.T) {
	var client = NewBambooClient(MockPayloadTransporter())
	permissions, err := client.ProjectService().ReadPermissions("PROJECT")
	assert.Nil(t, err)
	assert.NotNil(t, permissions)

	slices.SortFunc(permissions.Groups, func(a, b GroupPermission) int {
		return strings.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
	})
	slices.SortFunc(permissions.Users, func(a, b UserPermission) int {
		return strings.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
	})
	slices.SortFunc(permissions.Roles, func(a, b RolePermission) int {
		return strings.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
	})

	assert.Equal(t, "bamboo-admin", permissions.Groups[0].Name)
	assert.Equal(t, "bob", permissions.Users[0].Name)
	assert.Equal(t, "yunarta", permissions.Users[1].Name)
}

func TestProjectService_UpdateUserPermissions(t *testing.T) {
	var err error
	var ok bool
	var sent []transport.PayloadRequest

	transporter := MockPayloadTransporter()
	var client = NewBambooClient(transporter)

	err = client.ProjectService().UpdateUserPermissions("PROJECT", "yunarta", []string{"READ", "ADMINISTRATION"})
	assert.Nil(t, err)

	_, ok = transporter.SentRequests["DELETE:/rest/api/latest/permissions/project/PROJECT/users/yunarta"]
	assert.False(t, ok)

	sent, ok = transporter.SentRequests["PUT:/rest/api/latest/permissions/project/PROJECT/users/yunarta"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
}

func TestProjectService_UpdateUserPermissionsWithRemoval(t *testing.T) {
	var err error
	var ok bool
	var sent []transport.PayloadRequest

	transporter := MockPayloadTransporter()
	var client = NewBambooClient(transporter)

	err = client.ProjectService().UpdateUserPermissions("PROJECT", "yunarta", []string{"ADMINISTRATION"})
	assert.Nil(t, err)

	sent, ok = transporter.SentRequests["DELETE:/rest/api/latest/permissions/project/PROJECT/users/yunarta"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
	assert.Equal(t, "[\"READ\"]", string(sent[0].Payload.ContentMust()))

	sent, ok = transporter.SentRequests["PUT:/rest/api/latest/permissions/project/PROJECT/users/yunarta"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
}

func TestProjectService_UpdateGroupPermissions(t *testing.T) {
	var err error
	var ok bool
	var sent []transport.PayloadRequest

	transporter := MockPayloadTransporter()
	var client = NewBambooClient(transporter)

	err = client.ProjectService().UpdateGroupPermissions("PROJECT", "bamboo-admin", []string{"READ", "ADMINISTRATION"})
	assert.Nil(t, err)

	_, ok = transporter.SentRequests["DELETE:/rest/api/latest/permissions/project/PROJECT/groups/bamboo-admin"]
	assert.False(t, ok)

	sent, ok = transporter.SentRequests["PUT:/rest/api/latest/permissions/project/PROJECT/groups/bamboo-admin"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
}

func TestProjectService_UpdateGroupPermissionsWithRemoval(t *testing.T) {
	var err error
	var ok bool
	var sent []transport.PayloadRequest

	transporter := MockPayloadTransporter()
	var client = NewBambooClient(transporter)

	err = client.ProjectService().UpdateGroupPermissions("PROJECT", "bamboo-admin", []string{"ADMINISTRATION"})
	assert.Nil(t, err)

	sent, ok = transporter.SentRequests["DELETE:/rest/api/latest/permissions/project/PROJECT/groups/bamboo-admin"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
	assert.Equal(t, "[\"READ\"]", string(sent[0].Payload.ContentMust()))

	sent, ok = transporter.SentRequests["PUT:/rest/api/latest/permissions/project/PROJECT/groups/bamboo-admin"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
}

func TestProjectService_UpdateLoggedInRolePermissions(t *testing.T) {
	var err error
	var ok bool
	var sent []transport.PayloadRequest

	transporter := MockPayloadTransporter()
	var client = NewBambooClient(transporter)

	err = client.ProjectService().UpdateRolePermissions("PROJECT", "LOGGED_IN", []string{"READ", "ADMINISTRATION"})
	assert.Nil(t, err)

	_, ok = transporter.SentRequests["DELETE:/rest/api/latest/permissions/project/PROJECT/roles/LOGGED_IN"]
	assert.False(t, ok)

	sent, ok = transporter.SentRequests["PUT:/rest/api/latest/permissions/project/PROJECT/roles/LOGGED_IN"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
}

func TestProjectService_UpdateLoggedInRolePermissionsWithRemoval(t *testing.T) {
	var err error
	var ok bool
	var sent []transport.PayloadRequest

	transporter := MockPayloadTransporter()
	var client = NewBambooClient(transporter)

	err = client.ProjectService().UpdateRolePermissions("PROJECT", "LOGGED_IN", []string{"ADMINISTRATION"})
	assert.Nil(t, err)

	sent, ok = transporter.SentRequests["DELETE:/rest/api/latest/permissions/project/PROJECT/roles/LOGGED_IN"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
	assert.Equal(t, "[\"READ\"]", string(sent[0].Payload.ContentMust()))

	sent, ok = transporter.SentRequests["PUT:/rest/api/latest/permissions/project/PROJECT/roles/LOGGED_IN"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
}

func TestProjectService_UpdateAnonymousRolePermissions(t *testing.T) {
	var err error
	var ok bool
	var sent []transport.PayloadRequest

	transporter := MockPayloadTransporter()
	var client = NewBambooClient(transporter)

	err = client.ProjectService().UpdateRolePermissions("PROJECT2", "ANONYMOUS", []string{"READ"})
	assert.Nil(t, err)

	_, ok = transporter.SentRequests["DELETE:/rest/api/latest/permissions/project/PROJECT2/roles/ANONYMOUS"]
	assert.False(t, ok)

	sent, ok = transporter.SentRequests["PUT:/rest/api/latest/permissions/project/PROJECT2/roles/ANONYMOUS"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
	assert.Equal(t, "[\"READ\"]", string(sent[0].Payload.ContentMust()))
}

func TestProjectService_UpdateAnonymousRolePermissionsWithRemoval(t *testing.T) {
	var err error
	var ok bool
	var sent []transport.PayloadRequest

	transporter := MockPayloadTransporter()
	var client = NewBambooClient(transporter)

	err = client.ProjectService().UpdateRolePermissions("PROJECT", "ANONYMOUS", []string{})
	assert.Nil(t, err)

	sent, ok = transporter.SentRequests["DELETE:/rest/api/latest/permissions/project/PROJECT/roles/ANONYMOUS"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
	assert.Equal(t, "[\"READ\"]", string(sent[0].Payload.ContentMust()))

}
