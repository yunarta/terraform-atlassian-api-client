package bitbucket

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepositoryService_Create(t *testing.T) {
	var err error

	transporter := MockPayloadTransporter()
	var client = NewBitbucketClient(transporter)

	repository, err := client.RepositoryService().Create("A", CreateRepo{
		Name:        "name",
		Description: "description",
	})
	assert.Nil(t, err)
	assert.Equal(t, "name", repository.Name)
	assert.Equal(t, "description", repository.Description)
	assert.Equal(t, "slug", repository.Slug)
	assert.Equal(t, "A", repository.Project.Name)
	assert.Equal(t, "A", repository.Project.Key)
}

func TestRepositoryService_Read(t *testing.T) {
	var err error

	transporter := MockPayloadTransporter()
	var client = NewBitbucketClient(transporter)

	repository, err := client.RepositoryService().Read("A", "name")
	assert.Nil(t, err)
	assert.Equal(t, "name", repository.Name)
	assert.Equal(t, "description", repository.Description)
	assert.Equal(t, "slug", repository.Slug)
	assert.Equal(t, "A", repository.Project.Name)
	assert.Equal(t, "A", repository.Project.Key)
}

func TestRepositoryService_Update(t *testing.T) {
	var err error

	transporter := MockPayloadTransporter()
	var client = NewBitbucketClient(transporter)

	repository, err := client.RepositoryService().Update("A", "name", "new-description")
	assert.Nil(t, err)
	assert.Equal(t, "name", repository.Name)
	assert.Equal(t, "new-description", repository.Description)
	assert.Equal(t, "slug", repository.Slug)
	assert.Equal(t, "A", repository.Project.Name)
	assert.Equal(t, "A", repository.Project.Key)
}

func TestRepositoryService_Delete(t *testing.T) {
	var err error

	transporter := MockPayloadTransporter()
	var client = NewBitbucketClient(transporter)

	err = client.RepositoryService().Delete("A", "name")
	assert.Nil(t, err)
}

func TestRepositoryService_Initialize(t *testing.T) {
	var err error

	transporter := MockPayloadTransporter()
	var client = NewBitbucketClient(transporter)

	reply, err := client.RepositoryService().Initialize("A", "new", "# README")
	assert.Nil(t, err)
	assert.True(t, reply)

	reply, err = client.RepositoryService().Initialize("A", "exists", "# README")
	assert.Nil(t, err)
	assert.False(t, reply)
}
