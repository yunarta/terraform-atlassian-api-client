package bamboo

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yunarta/terraform-api-transport/transport"
	"testing"
)

func TestDeploymentService_Edge_Create(t *testing.T) {
	var client = NewBambooClient(EdgeTestMockPayloadTransporter())
	_, err := client.DeploymentService().Create(CreateDeployment{
		Name: "name",
		PlanKey: Key{
			Key: "BUILD-PLAN",
		},
		Description:  "description",
		PublicAccess: false,
	})
	assert.NotNil(t, err)

	var badRequestErr transport.BadRequestError
	assert.True(t, errors.As(err, &badRequestErr))
}

func TestDeploymentService_Edge_Read(t *testing.T) {
	var client = NewBambooClient(EdgeTestMockPayloadTransporter())
	deployment, err := client.DeploymentService().Read("dn")

	assert.NoError(t, err)
	fmt.Println(deployment)

	assert.Nil(t, deployment)
}

func TestDeploymentService_Edge_UpdateWithId(t *testing.T) {
	var client = NewBambooClient(EdgeTestMockPayloadTransporter())
	_, err := client.DeploymentService().UpdateWithId(2, UpdateDeployment{
		Name:        "name",
		PlanKey:     Key{Key: "BUILD-PLAN"},
		Description: "new-description",
	})
	assert.NotNil(t, err)

	var badRequestErr transport.BadRequestError
	assert.True(t, errors.As(err, &badRequestErr))
}

func TestDeploymentService_Edge_Delete(t *testing.T) {
	var client = NewBambooClient(EdgeTestMockPayloadTransporter())

	err := client.DeploymentService().Delete(2)

	assert.NotNil(t, err)

	var badRequestErr transport.BadRequestError
	assert.True(t, errors.As(err, &badRequestErr))
}

func TestDeploymentService_Edge_GetSpecRepositories(t *testing.T) {
	var client = NewBambooClient(EdgeTestMockPayloadTransporter())

	_, err := client.DeploymentService().GetSpecRepositories(2)
	assert.NotNil(t, err)

	var badRequestErr transport.BadRequestError
	assert.True(t, errors.As(err, &badRequestErr))
}

func TestDeploymentService_Edge_AddSpecRepositories(t *testing.T) {
	var client = NewBambooClient(EdgeTestMockPayloadTransporter())

	_, err := client.DeploymentService().AddSpecRepositories(2, 1)
	assert.NotNil(t, err)

	var badRequestErr transport.BadRequestError
	assert.True(t, errors.As(err, &badRequestErr))
}

func TestDeploymentService_Edge_RemoveSpecRepositories(t *testing.T) {
	var client = NewBambooClient(EdgeTestMockPayloadTransporter())

	err := client.DeploymentService().RemoveSpecRepositories(2, 1)
	assert.NotNil(t, err)

	var badRequestErr transport.BadRequestError
	assert.True(t, errors.As(err, &badRequestErr))
}
