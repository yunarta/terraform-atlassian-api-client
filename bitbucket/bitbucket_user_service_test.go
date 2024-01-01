package bitbucket

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService_FindUser(t *testing.T) {
	var client = NewBitbucketClient(MockPayloadTransporter())

	user, err := client.UserService().FindUser("yunarta")
	assert.Nil(t, err)

	assert.Equal(t, "yunarta", user.Name)

}

func TestUserService_FindGroup(t *testing.T) {
	var client = NewBitbucketClient(MockPayloadTransporter())

	user, err := client.UserService().FindGroup("bitbucket-admin")
	assert.Nil(t, err)

	assert.Equal(t, "bitbucket-admin", user.Name)

}
