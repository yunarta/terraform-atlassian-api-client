package bitbucket

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProjectService_Create(t *testing.T) {
	var client = NewBitbucketClient(MockPayloadTransporter())

	project, err := client.ProjectService().Create(CreateProject{
		Key:         "KEY",
		Name:        "NAME",
		Description: "DESCRIPTION",
	})
	assert.Nil(t, err)

	assert.NotNil(t, project)
	assert.Equal(t, "KEY", project.Key)
	assert.Equal(t, "NAME", project.Name)
	assert.Equal(t, "DESCRIPTION", project.Description)
}

func TestProjectService_Read(t *testing.T) {
	var client = NewBitbucketClient(MockPayloadTransporter())

	project, err := client.ProjectService().Read("KEY")
	assert.Nil(t, err)

	assert.NotNil(t, project)
	assert.Equal(t, "KEY", project.Key)
}

func TestProjectService_Update(t *testing.T) {
	var client = NewBitbucketClient(MockPayloadTransporter())

	updatedProject := ProjectUpdate{
		Name:        "NAME-UPDATED",
		Description: "DESCRIPTION-UPDATED",
	}

	project, err := client.ProjectService().Update("KEY", updatedProject)
	assert.Nil(t, err)

	assert.NotNil(t, project)
	assert.Equal(t, "NAME-UPDATED", project.Name)
	assert.Equal(t, "DESCRIPTION-UPDATED", project.Description)
}

func TestProjectService_Delete(t *testing.T) {
	var client = NewBitbucketClient(MockPayloadTransporter())

	err := client.ProjectService().Delete("KEY")
	assert.Nil(t, err)
}
