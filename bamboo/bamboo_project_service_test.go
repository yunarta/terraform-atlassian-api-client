package bamboo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProjectService_CreateProject(t *testing.T) {
	var err error

	transporter := MockPayloadTransporter()
	var client = NewBambooClient(transporter)

	project, err := client.ProjectService().Create(CreateProject{
		Name:         "name",
		Key:          "PROJECT",
		Description:  "description",
		PublicAccess: false,
	})

	assert.Nil(t, err)
	assert.Equal(t, "name", project.Name)
	assert.Equal(t, "PROJECT", project.Key)
	assert.Equal(t, "description", project.Description)
}

func TestProjectService_Read(t *testing.T) {
	var err error

	transporter := MockPayloadTransporter()
	var client = NewBambooClient(transporter)

	project, err := client.ProjectService().Read("PROJECT")

	assert.Nil(t, err)
	assert.Equal(t, "name", project.Name)
	assert.Equal(t, "PROJECT", project.Key)
	assert.Equal(t, "description", project.Description)
}

func TestProjectService_Update(t *testing.T) {
	var err error

	transporter := MockPayloadTransporter()
	var client = NewBambooClient(transporter)

	project, err := client.ProjectService().Update("PROJECT", UpdateProject{
		Name:        "New Name",
		Description: "New Description",
	})

	assert.Nil(t, err)
	assert.Equal(t, "New Name", project.Name)
	assert.Equal(t, "New Description", project.Description)
}

func TestProjectService_Delete(t *testing.T) {
	var err error

	transporter := MockPayloadTransporter()
	var client = NewBambooClient(transporter)

	err = client.ProjectService().Delete("PROJECT")

	assert.Nil(t, err)
}

func TestProjectService_ReadPlan(t *testing.T) {
	var err error

	transporter := MockPayloadTransporter()
	var client = NewBambooClient(transporter)
	plan, err := client.ProjectService().ReadPlan("PROJECT-PLAN")

	assert.NoError(t, err)
	assert.NotNil(t, plan)

	assert.Equal(t, "PROJECT", plan.ProjectKey)
	assert.Equal(t, "PROJECT", plan.ProjectName)
	assert.Equal(t, "description", plan.Description)
	assert.Equal(t, "PLAN", plan.Key)
	assert.Equal(t, "short name", plan.ShortName)
	assert.Equal(t, "name", plan.Name)
}

func TestProjectService_GetSpecRepositories(t *testing.T) {
	var client = NewBambooClient(MockPayloadTransporter())

	repository, err := client.ProjectService().GetSpecRepositories("PROJECT")
	assert.NoError(t, err)

	assert.NotNil(t, repository)
	assert.Equal(t, 1, repository[0].ID)
	assert.Equal(t, "repository-1", repository[0].Name)
	assert.Equal(t, false, repository[0].RssEnabled)
}

func TestProjectService_AddSpecRepositories(t *testing.T) {
	var client = NewBambooClient(MockPayloadTransporter())

	repository, err := client.ProjectService().AddSpecRepositories("PROJECT", 1)
	assert.NoError(t, err)

	assert.NotNil(t, repository)
	assert.Equal(t, 1, repository.ID)
	assert.Equal(t, "repository-1", repository.Name)
	assert.Equal(t, false, repository.RssEnabled)
}

func TestProjectService_RemoveSpecRepositories(t *testing.T) {
	var client = NewBambooClient(MockPayloadTransporter())

	err := client.ProjectService().RemoveSpecRepositories("PROJECT", 1)
	assert.NoError(t, err)
}
