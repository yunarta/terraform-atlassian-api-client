package cloud

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadUser(t *testing.T) {
	var err error

	var client = NewJiraClient(MockPayloadTransporter())

	success, err := client.ActorService().ReadUser("yunarta.kartawahyudi@gmail.com")
	assert.Nil(t, err)
	assert.NotNil(t, success)
}

func TestReadGroup(t *testing.T) {
	var err error

	var client = NewJiraClient(MockPayloadTransporter())

	success, err := client.ActorService().ReadGroup("jira-admins-mobilesolutionworks")
	assert.Nil(t, err)
	assert.NotNil(t, success)
}
