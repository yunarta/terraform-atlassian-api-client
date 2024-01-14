package cloud

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestActorService_ReadUser(t *testing.T) {
	var err error

	var client = NewJiraClient(MockPayloadTransporter())

	success, err := client.ActorService().ReadUser("yunarta.kartawahyudi@gmail.com")
	assert.Nil(t, err)
	assert.NotNil(t, success)
}

func TestActorService_ReadGroup(t *testing.T) {
	var err error

	var client = NewJiraClient(MockPayloadTransporter())

	success, err := client.ActorService().ReadGroup("jira-admins-mobilesolutionworks")
	assert.Nil(t, err)
	assert.NotNil(t, success)
}

func TestActorService_BulkGetUsers(t *testing.T) {
	var err error

	var client = NewJiraClient(MockPayloadTransporter())

	users, err := client.ActorService().BulkGetUsers([]string{"557058:32b276cf-1a9f-45ae-b3f5-f850bc24f1b9"})
	assert.Nil(t, err)
	assert.NotNil(t, users)

	assert.Equal(t, "yunarta.kartawahyudi@gmail.com", users[0].EmailAddress)
	assert.Equal(t, "Yunarta Kartawahyudi", users[0].DisplayName)
}

func TestActorService_BulkGetGroupsById(t *testing.T) {
	var err error

	var client = NewJiraClient(MockPayloadTransporter())

	users, err := client.ActorService().BulkGetGroupsById([]string{"94af0e5e-018a-422d-bffa-e41fc5b71d29"})
	assert.Nil(t, err)
	assert.NotNil(t, users)

	assert.Equal(t, "site-admins", users[0].Name)
	assert.Equal(t, "94af0e5e-018a-422d-bffa-e41fc5b71d29", users[0].GroupId)
}

func TestActorServiceBulkGetGroupsByName(t *testing.T) {
	var err error

	var client = NewJiraClient(MockPayloadTransporter())

	users, err := client.ActorService().BulkGetGroupsByName([]string{"site-admins"})
	assert.Nil(t, err)
	assert.NotNil(t, users)

	assert.Equal(t, "site-admins", users[0].Name)
	assert.Equal(t, "94af0e5e-018a-422d-bffa-e41fc5b71d29", users[0].GroupId)
}
