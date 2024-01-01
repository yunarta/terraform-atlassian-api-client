package bamboo

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/yunarta/terraform-api-transport/transport"
	"testing"
)

func TestProjectService_Edge_ReadPermissions(t *testing.T) {
	var client = NewBambooClient(EdgeTestMockPayloadTransporter())
	_, err := client.ProjectService().ReadPermissions("PROJECT")
	assert.NotNil(t, err)

	var badRequestErr transport.BadRequestError
	assert.True(t, errors.As(err, &badRequestErr))
}

func TestProjectService_Edge_UpdateUserPermissions(t *testing.T) {
	var err error
	var client = NewBambooClient(EdgeTestMockPayloadTransporter())

	err = client.ProjectService().UpdateUserPermissions("PROJECT", "yunarta", []string{"READ", "ADMINISTRATION"})
	assert.NotNil(t, err)

	var badRequestErr transport.BadRequestError
	assert.True(t, errors.As(err, &badRequestErr))
}

func TestProjectService_Edge_UpdateGroupPermissions(t *testing.T) {
	var err error

	var client = NewBambooClient(EdgeTestMockPayloadTransporter())

	err = client.ProjectService().UpdateGroupPermissions("PROJECT", "bamboo-admin", []string{"READ", "ADMINISTRATION"})
	assert.NotNil(t, err)

	var badRequestErr transport.BadRequestError
	assert.True(t, errors.As(err, &badRequestErr))
}

func TestProjectService_Edge_UpdateLoggedInRolePermissions(t *testing.T) {
	var err error

	var client = NewBambooClient(EdgeTestMockPayloadTransporter())

	err = client.ProjectService().UpdateRolePermissions("PROJECT", "LOGGED_IN", []string{"READ", "ADMINISTRATION"})
	assert.NotNil(t, err)

	var badRequestErr transport.BadRequestError
	assert.True(t, errors.As(err, &badRequestErr))
}
