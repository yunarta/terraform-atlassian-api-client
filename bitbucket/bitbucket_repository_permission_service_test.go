package bitbucket

import (
	"github.com/stretchr/testify/assert"
	"github.com/yunarta/terraform-api-transport/transport"
	"testing"
)

func TestRepositoryService_ReadPermissions(t *testing.T) {
	var client = NewBitbucketClient(MockPayloadTransporter())

	permissions, err := client.RepositoryService().ReadPermissions("A", "r")
	assert.Nil(t, err)

	assert.Equal(t, "yunarta", permissions.Users[0].Owner.Name)
	assert.Equal(t, "bitbucket-admin", permissions.Groups[0].Owner.Name)
}

func TestRepositoryService_UpdateUserPermission(t *testing.T) {
	var err error
	var ok bool
	var sent []transport.PayloadRequest

	transporter := MockPayloadTransporter()
	var client = NewBitbucketClient(transporter)

	err = client.RepositoryService().UpdateUserPermission("A", "r", "yunarta", "REPO_ADMIN")
	assert.Nil(t, err)

	_, ok = transporter.SentRequests["DELETE:/rest/api/latest/projects/A/repos/r/permissions/users?name=yunarta"]
	assert.False(t, ok)

	sent, ok = transporter.SentRequests["PUT:/rest/api/latest/projects/A/repos/r/permissions/users?name=yunarta&permission=REPO_ADMIN"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
}

func TestRepositoryService_UpdateUserPermissionRemoval(t *testing.T) {
	var err error
	var ok bool

	transporter := MockPayloadTransporter()
	var client = NewBitbucketClient(transporter)

	err = client.RepositoryService().UpdateUserPermission("A", "r", "yunarta", "")
	assert.Nil(t, err)

	_, ok = transporter.SentRequests["DELETE:/rest/api/latest/projects/A/repos/r/permissions/users?name=yunarta"]
	assert.True(t, ok)
}

func TestRepositoryService_UpdateUserPermissionNew(t *testing.T) {
	var err error
	var ok bool
	var sent []transport.PayloadRequest

	transporter := MockPayloadTransporter()
	var client = NewBitbucketClient(transporter)

	err = client.RepositoryService().UpdateUserPermission("B", "r", "yunarta", "REPO_ADMIN")
	assert.Nil(t, err)

	sent, ok = transporter.SentRequests["PUT:/rest/api/latest/projects/B/repos/r/permissions/users?name=yunarta&permission=REPO_ADMIN"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
}

func TestRepositoryService_UpdateGroupPermission(t *testing.T) {
	var err error
	var ok bool
	var sent []transport.PayloadRequest

	transporter := MockPayloadTransporter()
	var client = NewBitbucketClient(transporter)

	err = client.RepositoryService().UpdateGroupPermission("A", "r", "bitbucket-admin", "REPO_ADMIN")
	assert.Nil(t, err)

	_, ok = transporter.SentRequests["DELETE:/rest/api/latest/projects/A/repos/r/permissions/groups?name=bitbucket-admin"]
	assert.False(t, ok)

	sent, ok = transporter.SentRequests["PUT:/rest/api/latest/projects/A/repos/r/permissions/groups?name=bitbucket-admin&permission=REPO_ADMIN"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
}

func TestRepositoryService_UpdateGroupPermissionRemoval(t *testing.T) {
	var err error
	var ok bool

	transporter := MockPayloadTransporter()
	var client = NewBitbucketClient(transporter)

	err = client.RepositoryService().UpdateGroupPermission("A", "r", "bitbucket-admin", "")
	assert.Nil(t, err)

	_, ok = transporter.SentRequests["DELETE:/rest/api/latest/projects/A/repos/r/permissions/groups?name=bitbucket-admin"]
	assert.True(t, ok)
}

func TestRepositoryService_UpdateGroupPermissionNew(t *testing.T) {
	var err error
	var ok bool
	var sent []transport.PayloadRequest

	transporter := MockPayloadTransporter()
	var client = NewBitbucketClient(transporter)

	err = client.RepositoryService().UpdateGroupPermission("B", "r", "bitbucket-admin", "REPO_ADMIN")
	assert.Nil(t, err)

	sent, ok = transporter.SentRequests["PUT:/rest/api/latest/projects/B/repos/r/permissions/groups?name=bitbucket-admin&permission=REPO_ADMIN"]
	assert.True(t, ok)
	assert.Equal(t, 1, len(sent))
}
