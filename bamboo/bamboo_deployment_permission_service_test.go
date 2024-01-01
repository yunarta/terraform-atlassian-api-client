package bamboo

import (
	"github.com/stretchr/testify/assert"
	"github.com/yunarta/terraform-api-transport/transport"
	"slices"
	"strings"
	"testing"
)

func TestDeploymentService_ReadPermissions(t *testing.T) {
	var client = NewBambooClient(MockPayloadTransporter())
	permissions, err := client.DeploymentService().ReadPermissions(0)
	assert.Nil(t, err)

	slices.SortFunc(permissions.Groups, func(a, b GroupPermission) int {
		return strings.Compare(a.Name, b.Name)
	})
	slices.SortFunc(permissions.Users, func(a, b UserPermission) int {
		return strings.Compare(a.Name, b.Name)
	})
	slices.SortFunc(permissions.Roles, func(a, b RolePermission) int {
		return strings.Compare(a.Name, b.Name)
	})

	assert.NotNil(t, permissions)
	assert.Contains(t, permissions.Groups[0].Name, "bamboo-admin")
	assert.Contains(t, permissions.Users[0].Name, "yunarta")
	assert.Contains(t, permissions.Roles[0].Name, "ANONYMOUS")
	assert.Contains(t, permissions.Roles[1].Name, "LOGGED_IN")
}

func TestDeploymentService_UpdateUserPermissions(t *testing.T) {
	var err error
	var ok bool
	var sent []transport.PayloadRequest

	transporter := MockPayloadTransporter()
	var client = NewBambooClient(transporter)

	err = client.DeploymentService().UpdateUserPermissions(0, "yunarta", []string{"READ", "WRITE"})
	assert.Nil(t, err)

	_, ok = transporter.SentRequests["DELETE:/rest/api/latest/permissions/deployment/0/users/yunarta"]
	assert.False(t, ok)

	sent, ok = transporter.SentRequests["PUT:/rest/api/latest/permissions/deployment/0/users/yunarta"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
}

func TestDeploymentService_UpdateUserPermissionsWithRemoval(t *testing.T) {
	var err error
	var ok bool
	var sent []transport.PayloadRequest

	transporter := MockPayloadTransporter()
	var client = NewBambooClient(transporter)

	err = client.DeploymentService().UpdateUserPermissions(0, "yunarta", []string{"WRITE"})
	assert.Nil(t, err)

	sent, ok = transporter.SentRequests["DELETE:/rest/api/latest/permissions/deployment/0/users/yunarta"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
	assert.Equal(t, "[\"READ\"]", string(sent[0].Payload.ContentMust()))

	sent, ok = transporter.SentRequests["PUT:/rest/api/latest/permissions/deployment/0/users/yunarta"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
}

func TestDeploymentService_UpdateGroupPermissions(t *testing.T) {
	var err error
	var ok bool
	var sent []transport.PayloadRequest

	transporter := MockPayloadTransporter()
	var client = NewBambooClient(transporter)

	err = client.DeploymentService().UpdateGroupPermissions(0, "bamboo-admin", []string{"READ", "WRITE"})
	assert.Nil(t, err)

	_, ok = transporter.SentRequests["DELETE:/rest/api/latest/permissions/deployment/0/groups/bamboo-admin"]
	assert.False(t, ok)

	sent, ok = transporter.SentRequests["PUT:/rest/api/latest/permissions/deployment/0/groups/bamboo-admin"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
}

func TestDeploymentService_UpdateGroupPermissionsWithRemoval(t *testing.T) {
	var err error
	var ok bool
	var sent []transport.PayloadRequest

	transporter := MockPayloadTransporter()
	var client = NewBambooClient(transporter)

	err = client.DeploymentService().UpdateGroupPermissions(0, "bamboo-admin", []string{"WRITE"})
	assert.Nil(t, err)

	sent, ok = transporter.SentRequests["DELETE:/rest/api/latest/permissions/deployment/0/groups/bamboo-admin"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
	assert.Equal(t, "[\"READ\"]", string(sent[0].Payload.ContentMust()))

	sent, ok = transporter.SentRequests["PUT:/rest/api/latest/permissions/deployment/0/groups/bamboo-admin"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
}

func TestDeploymentService_UpdateRolePermissions(t *testing.T) {
	var err error
	var ok bool
	var sent []transport.PayloadRequest

	transporter := MockPayloadTransporter()
	var client = NewBambooClient(transporter)

	err = client.DeploymentService().UpdateRolePermissions(0, "LOGGED_IN", []string{"READ", "WRITE"})
	assert.Nil(t, err)

	_, ok = transporter.SentRequests["DELETE:/rest/api/latest/permissions/deployment/0/roles/LOGGED_IN"]
	assert.False(t, ok)

	sent, ok = transporter.SentRequests["PUT:/rest/api/latest/permissions/deployment/0/roles/LOGGED_IN"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
}

func TestDeploymentService_UpdateRolePermissionsWithRemoval(t *testing.T) {
	var err error
	var ok bool
	var sent []transport.PayloadRequest

	transporter := MockPayloadTransporter()
	var client = NewBambooClient(transporter)

	err = client.DeploymentService().UpdateRolePermissions(0, "LOGGED_IN", []string{"WRITE"})
	assert.Nil(t, err)

	sent, ok = transporter.SentRequests["DELETE:/rest/api/latest/permissions/deployment/0/roles/LOGGED_IN"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
	assert.Equal(t, "[\"READ\"]", string(sent[0].Payload.ContentMust()))

	sent, ok = transporter.SentRequests["PUT:/rest/api/latest/permissions/deployment/0/roles/LOGGED_IN"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
}

func TestDeploymentService_UpdateUsersPermissionsNew(t *testing.T) {
	var err error
	var ok bool
	var sent []transport.PayloadRequest

	transporter := MockPayloadTransporter()
	var client = NewBambooClient(transporter)

	err = client.DeploymentService().UpdateUserPermissions(1, "yunarta", []string{"WRITE"})
	assert.Nil(t, err)

	sent, ok = transporter.SentRequests["PUT:/rest/api/latest/permissions/deployment/1/users/yunarta"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
}
