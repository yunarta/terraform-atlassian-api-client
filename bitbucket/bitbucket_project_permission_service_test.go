package bitbucket

import (
	"github.com/stretchr/testify/assert"
	"github.com/yunarta/terraform-api-transport/transport"
	"testing"
)

func TestProjectService_ReadPermissions(t *testing.T) {
	var client = NewBitbucketClient(MockPayloadTransporter())

	permissions, err := client.ProjectService().ReadPermissions("A")
	assert.Nil(t, err)

	assert.Equal(t, "yunarta", permissions.Users[0].Owner.Name)
	assert.Equal(t, "bitbucket-admin", permissions.Groups[0].Owner.Name)
}

func TestProjectService_UpdateUserPermission(t *testing.T) {
	var err error
	var ok bool
	var sent []transport.PayloadRequest

	transporter := MockPayloadTransporter()
	var client = NewBitbucketClient(transporter)

	err = client.ProjectService().UpdateUserPermission("A", "yunarta", "PROJECT_ADMIN")
	assert.Nil(t, err)

	_, ok = transporter.SentRequests["DELETE:/rest/api/latest/projects/A/permissions/users?name=yunarta"]
	assert.False(t, ok)

	sent, ok = transporter.SentRequests["PUT:/rest/api/latest/projects/A/permissions/users?name=yunarta&permission=PROJECT_ADMIN"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
}

func TestProjectService_UpdateUserPermissionRemoval(t *testing.T) {
	var err error
	var ok bool

	transporter := MockPayloadTransporter()
	var client = NewBitbucketClient(transporter)

	err = client.ProjectService().UpdateUserPermission("A", "yunarta", "")
	assert.Nil(t, err)

	_, ok = transporter.SentRequests["DELETE:/rest/api/latest/projects/A/permissions/users?name=yunarta"]
	assert.True(t, ok)
}

func TestProjectService_UpdateUserPermissionNew(t *testing.T) {
	var err error
	var ok bool
	var sent []transport.PayloadRequest

	transporter := MockPayloadTransporter()
	var client = NewBitbucketClient(transporter)

	err = client.ProjectService().UpdateUserPermission("B", "yunarta", "PROJECT_ADMIN")
	assert.Nil(t, err)

	sent, ok = transporter.SentRequests["PUT:/rest/api/latest/projects/B/permissions/users?name=yunarta&permission=PROJECT_ADMIN"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
}

func TestProjectService_UpdateGroupPermission(t *testing.T) {
	var err error
	var ok bool
	var sent []transport.PayloadRequest

	transporter := MockPayloadTransporter()
	var client = NewBitbucketClient(transporter)

	err = client.ProjectService().UpdateGroupPermission("A", "bitbucket-admin", "PROJECT_ADMIN")
	assert.Nil(t, err)

	_, ok = transporter.SentRequests["DELETE:/rest/api/latest/projects/A/permissions/groups?name=bitbucket-admin"]
	assert.False(t, ok)

	sent, ok = transporter.SentRequests["PUT:/rest/api/latest/projects/A/permissions/groups?name=bitbucket-admin&permission=PROJECT_ADMIN"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
}

func TestProjectService_UpdateGroupPermissionRemoval(t *testing.T) {
	var err error
	var ok bool

	transporter := MockPayloadTransporter()
	var client = NewBitbucketClient(transporter)

	err = client.ProjectService().UpdateGroupPermission("A", "bitbucket-admin", "")
	assert.Nil(t, err)

	_, ok = transporter.SentRequests["DELETE:/rest/api/latest/projects/A/permissions/groups?name=bitbucket-admin"]
	assert.True(t, ok)
}

func TestProjectService_UpdateGroupPermissionNew(t *testing.T) {
	var err error
	var ok bool
	var sent []transport.PayloadRequest

	transporter := MockPayloadTransporter()
	var client = NewBitbucketClient(transporter)

	err = client.ProjectService().UpdateGroupPermission("B", "bitbucket-admin", "PROJECT_ADMIN")
	assert.Nil(t, err)

	sent, ok = transporter.SentRequests["PUT:/rest/api/latest/projects/B/permissions/groups?name=bitbucket-admin&permission=PROJECT_ADMIN"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
}
