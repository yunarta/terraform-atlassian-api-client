package cloud

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestActorLookupService_FindUser_WithoutRegisterAccountIds(t *testing.T) {
	client := NewJiraClient(MockPayloadTransporter())

	user := client.ActorLookupService().FindUser("yunarta.kartawahyudi@gmail.com")

	assert.Equal(t, "557058:32b276cf-1a9f-45ae-b3f5-f850bc24f1b9", user.AccountID)
	assert.Equal(t, "yunarta.kartawahyudi@gmail.com", user.EmailAddress)
}

func TestActorLookupService_FindUser_WithRegisterAccountIds(t *testing.T) {
	client := NewJiraClient(MockPayloadTransporter())

	client.ActorLookupService().RegisterAccountIds("557058:32b276cf-1a9f-45ae-b3f5-f850bc24f1b9")
	user := client.ActorLookupService().FindUser("yunarta.kartawahyudi@gmail.com")

	assert.Equal(t, "557058:32b276cf-1a9f-45ae-b3f5-f850bc24f1b9", user.AccountID)
	assert.Equal(t, "yunarta.kartawahyudi@gmail.com", user.EmailAddress)
}

func TestActorLookupService_FindUser_WithRegisterUsernames(t *testing.T) {
	client := NewJiraClient(MockPayloadTransporter())

	client.ActorLookupService().RegisterUsernames("yunarta.kartawahyudi@gmail.com")
	user := client.ActorLookupService().FindUser("yunarta.kartawahyudi@gmail.com")

	assert.Equal(t, "557058:32b276cf-1a9f-45ae-b3f5-f850bc24f1b9", user.AccountID)
	assert.Equal(t, "yunarta.kartawahyudi@gmail.com", user.EmailAddress)
}

func TestActorLookupService_FindUserById_WithoutRegisterAccountIds(t *testing.T) {
	client := NewJiraClient(MockPayloadTransporter())

	user := client.ActorLookupService().FindUserById("557058:32b276cf-1a9f-45ae-b3f5-f850bc24f1b9")

	assert.Equal(t, "557058:32b276cf-1a9f-45ae-b3f5-f850bc24f1b9", user.AccountID)
	assert.Equal(t, "yunarta.kartawahyudi@gmail.com", user.EmailAddress)
}

func TestActorLookupService_FindUserById_WithRegisterAccountIds(t *testing.T) {
	client := NewJiraClient(MockPayloadTransporter())

	client.ActorLookupService().RegisterAccountIds("557058:32b276cf-1a9f-45ae-b3f5-f850bc24f1b9")
	user := client.ActorLookupService().FindUserById("557058:32b276cf-1a9f-45ae-b3f5-f850bc24f1b9")

	assert.Equal(t, "557058:32b276cf-1a9f-45ae-b3f5-f850bc24f1b9", user.AccountID)
	assert.Equal(t, "yunarta.kartawahyudi@gmail.com", user.EmailAddress)
}

func TestActorLookupService_FindGroup_WithoutRegisterAccountIds(t *testing.T) {
	client := NewJiraClient(MockPayloadTransporter())

	user := client.ActorLookupService().FindGroup("jira-admins-mobilesolutionworks")

	assert.Equal(t, "2eeab65d-ab39-4eee-965b-fd2a8bed9d12", user.GroupId)
	assert.Equal(t, "jira-admins-mobilesolutionworks", user.Name)
}

func TestActorLookupService_FindGroup_WithRegisterAccountIds(t *testing.T) {
	client := NewJiraClient(MockPayloadTransporter())

	client.ActorLookupService().RegisterGroupIds("2eeab65d-ab39-4eee-965b-fd2a8bed9d12")
	user := client.ActorLookupService().FindGroup("jira-admins-mobilesolutionworks")

	assert.Equal(t, "2eeab65d-ab39-4eee-965b-fd2a8bed9d12", user.GroupId)
	assert.Equal(t, "jira-admins-mobilesolutionworks", user.Name)
}

func TestActorLookupService_FindGroup_WithRegisterUsernames(t *testing.T) {
	client := NewJiraClient(MockPayloadTransporter())

	client.ActorLookupService().RegisterGroupNames("jira-admins-mobilesolutionworks")
	group := client.ActorLookupService().FindGroup("jira-admins-mobilesolutionworks")

	assert.Equal(t, "2eeab65d-ab39-4eee-965b-fd2a8bed9d12", group.GroupId)
	assert.Equal(t, "jira-admins-mobilesolutionworks", group.Name)
}

func TestActorLookupService_FindGroupById_WithoutRegisterAccountIds(t *testing.T) {
	client := NewJiraClient(MockPayloadTransporter())

	group := client.ActorLookupService().FindGroupById("2eeab65d-ab39-4eee-965b-fd2a8bed9d12")

	assert.Equal(t, "2eeab65d-ab39-4eee-965b-fd2a8bed9d12", group.GroupId)
	assert.Equal(t, "jira-admins-mobilesolutionworks", group.Name)
}

func TestActorLookupService_FindGroupById_WithRegisterAccountIds(t *testing.T) {
	client := NewJiraClient(MockPayloadTransporter())

	client.ActorLookupService().RegisterGroupIds("2eeab65d-ab39-4eee-965b-fd2a8bed9d12")
	group := client.ActorLookupService().FindGroupById("2eeab65d-ab39-4eee-965b-fd2a8bed9d12")

	assert.Equal(t, "2eeab65d-ab39-4eee-965b-fd2a8bed9d12", group.GroupId)
	assert.Equal(t, "jira-admins-mobilesolutionworks", group.Name)
}
