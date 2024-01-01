package bamboo

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/yunarta/terraform-api-transport/transport"
	"testing"
)

func TestDeploymentService_Edge_ReadPermissions(t *testing.T) {
	var client = NewBambooClient(EdgeTestMockPayloadTransporter())
	_, err := client.DeploymentService().ReadPermissions(0)
	assert.NotNil(t, err)

	var badRequestErr transport.BadRequestError
	assert.True(t, errors.As(err, &badRequestErr))
}

func TestDeploymentService_Edge_UpdateUserPermissions(t *testing.T) {
	var client = NewBambooClient(EdgeTestMockPayloadTransporter())
	err := client.DeploymentService().UpdateUserPermissions(0, "yunarta", []string{"READ", "WRITE"})
	assert.NotNil(t, err)

	var badRequestErr transport.BadRequestError
	assert.True(t, errors.As(err, &badRequestErr))
}

func TestDeploymentService_Edge_UpdateGroupPermissions(t *testing.T) {
	var err error

	transporter := EdgeTestMockPayloadTransporter()
	var client = NewBambooClient(transporter)

	err = client.DeploymentService().UpdateGroupPermissions(0, "bamboo-admin", []string{"READ", "WRITE"})
	assert.NotNil(t, err)

	var badRequestErr transport.BadRequestError
	assert.True(t, errors.As(err, &badRequestErr))
}

func TestDeploymentService_Edge_UpdateRolePermissions(t *testing.T) {
	var err error

	transporter := EdgeTestMockPayloadTransporter()
	var client = NewBambooClient(transporter)

	err = client.DeploymentService().UpdateRolePermissions(0, "LOGGED_IN", []string{"READ", "WRITE"})
	assert.NotNil(t, err)

	var badRequestErr transport.BadRequestError
	assert.True(t, errors.As(err, &badRequestErr))
}
