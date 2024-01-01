package bamboo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService_CurrentUser(t *testing.T) {
	var client = NewBambooClient(MockPayloadTransporter())

	user, err := client.UserService().CurrentUser()
	assert.Nil(t, err)

	assert.Equal(t, "yunarta", user.Name)
	assert.Equal(t, "Yunarta", user.FullName)
	assert.Equal(t, "yunarta.kartawahyudi@gmail.com", user.Email)
}

func TestUserService_FindUser(t *testing.T) {
	var client = NewBambooClient(MockPayloadTransporter())

	user, err := client.UserService().FindUser("yunarta")
	assert.Nil(t, err)

	assert.Equal(t, "yunarta", user.Name)

}

func TestUserService_FindGroup(t *testing.T) {
	var client = NewBambooClient(MockPayloadTransporter())

	user, err := client.UserService().FindGroup("bamboo-admin")
	assert.Nil(t, err)

	assert.Equal(t, "bamboo-admin", user.Name)

}
