package bamboo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeploymentService_Create(t *testing.T) {
	var client = NewBambooClient(MockPayloadTransporter())
	deployment, err := client.DeploymentService().Create(CreateDeployment{
		Name: "name",
		PlanKey: Key{
			Key: "BUILD-PLAN",
		},
		Description:  "description",
		PublicAccess: false,
	})

	assert.NoError(t, err)
	assert.Equal(t, "name", deployment.Name)
	assert.Equal(t, "BUILD-PLAN", deployment.PlanKey.Key)
	assert.Equal(t, "description", deployment.Description)
	assert.Empty(t, deployment.Environments)
	assert.Equal(t, false, deployment.RepositorySpecsManaged)
}

func TestDeploymentService_Read(t *testing.T) {
	var client = NewBambooClient(MockPayloadTransporter())
	deployment, err := client.DeploymentService().Read("dn")

	assert.NoError(t, err)
	fmt.Println(deployment)

	assert.Equal(t, 2, deployment.ID)
	assert.Equal(t, "dn", deployment.Name)
	assert.Equal(t, "BUILD-PLAN", deployment.PlanKey.Key)
	assert.Equal(t, "description", deployment.Description)
}

func TestDeploymentService_UpdateWithId(t *testing.T) {
	var client = NewBambooClient(MockPayloadTransporter())
	deployment, err := client.DeploymentService().ReadWithId(2)

	assert.NoError(t, err)

	deployment, err = client.DeploymentService().UpdateWithId(2, UpdateDeployment{
		Name:        deployment.Name,
		PlanKey:     Key{Key: deployment.PlanKey.Key},
		Description: "new-description",
	})

	assert.NoError(t, err)
	assert.Equal(t, "new-description", deployment.Description)
}

func TestDeploymentService_Delete(t *testing.T) {
	var client = NewBambooClient(MockPayloadTransporter())

	err := client.DeploymentService().Delete(2)

	assert.NoError(t, err)
}

func TestDeploymentService_GetSpecRepositories(t *testing.T) {
	var client = NewBambooClient(MockPayloadTransporter())

	repository, err := client.DeploymentService().GetSpecRepositories(2)
	assert.NoError(t, err)

	assert.NotNil(t, repository)
	assert.Equal(t, 1, repository[0].ID)
	assert.Equal(t, "repository-1", repository[0].Name)
	assert.Equal(t, false, repository[0].RssEnabled)
}

func TestDeploymentService_AddSpecRepositories(t *testing.T) {
	var client = NewBambooClient(MockPayloadTransporter())

	repository, err := client.DeploymentService().AddSpecRepositories(2, 1)
	assert.NoError(t, err)

	assert.NotNil(t, repository)
	assert.Equal(t, 1, repository.ID)
	assert.Equal(t, "repository-1", repository.Name)
	assert.Equal(t, false, repository.RssEnabled)
}

func TestDeploymentService_RemoveSpecRepositories(t *testing.T) {
	var client = NewBambooClient(MockPayloadTransporter())

	err := client.DeploymentService().RemoveSpecRepositories(2, 1)
	assert.NoError(t, err)
}
