package bamboo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepositoryService_Create(t *testing.T) {
	var client = NewBambooClient(MockPayloadTransporter())

	reply, err := client.RepositoryService().Create(CreateRepository{
		Name:           "create",
		ProjectKey:     "containers",
		RepositorySlug: "ubuntu",
		ServerId:       "55f10dce-f2c6-3885-b21f-2742d85ea4ea",
		ServerName:     "Bitbucket",
		CloneUrl:       "ssh://git@bitbucket.funf:7999/containers/b1.git",
	})

	assert.Nil(t, err)
	assert.Equal(t, 1, reply)
}

func TestRepositoryService_Read(t *testing.T) {
	var client = NewBambooClient(MockPayloadTransporter())

	repository, err := client.RepositoryService().Read("oraclelinux")

	assert.Nil(t, err)
	assert.Equal(t, "oraclelinux", repository.Name)
	assert.True(t, repository.RssEnabled)
}

func TestRepositoryService_Update(t *testing.T) {
	var client = NewBambooClient(MockPayloadTransporter())

	repository, err := client.RepositoryService().Read("oraclelinux")
	assert.Nil(t, err)

	err = client.RepositoryService().Update(repository.ID, true)
	assert.Nil(t, err)
}

func TestRepositoryService_ReadAccessor(t *testing.T) {
	var client = NewBambooClient(MockPayloadTransporter())

	repository, err := client.RepositoryService().Read("oraclelinux")
	assert.Nil(t, err)

	accessor, err := client.RepositoryService().ReadAccessor(repository.ID)
	assert.Nil(t, err)
	assert.Equal(t, "ubuntu", accessor[0].Name)
	assert.True(t, accessor[0].RssEnabled)
}

func TestRepositoryService_AddAccessor(t *testing.T) {
	var client = NewBambooClient(MockPayloadTransporter())

	repository, err := client.RepositoryService().Read("oraclelinux")
	assert.Nil(t, err)

	accessor, err := client.RepositoryService().Read("ubuntu")
	assert.Nil(t, err)

	_, err = client.RepositoryService().AddAccessor(repository.ID, accessor.ID)
	assert.Nil(t, err)
}

func TestRepositoryService_RemoveAccessor(t *testing.T) {
	var client = NewBambooClient(MockPayloadTransporter())

	repository, err := client.RepositoryService().Read("oraclelinux")
	assert.Nil(t, err)

	accessor, err := client.RepositoryService().Read("ubuntu")
	assert.Nil(t, err)

	err = client.RepositoryService().RemoveAccessor(repository.ID, accessor.ID)
	assert.Nil(t, err)
}
