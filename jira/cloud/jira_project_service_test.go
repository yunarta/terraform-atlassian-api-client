package cloud

import (
	"github.com/stretchr/testify/assert"
	"github.com/yunarta/terraform-atlassian-api-client/jira"
)

import "testing"

func TestProjectService_Create(t *testing.T) {
	var err error

	var client = NewJiraClient(MockPayloadTransporter())

	reply, err := client.ProjectService().Create(jira.CreateProject{
		Key:            "PROJECT",
		Name:           "PROJECT",
		ProjectTypeKey: "software",
		LeadAccountId:  "557058:32b276cf-1a9f-45ae-b3f5-f850bc24f1b9",
	})
	assert.Nil(t, err)
	assert.Equal(t, "PROJECT", reply.Name)
}

func TestProjectService_Clone(t *testing.T) {
	var err error

	var client = NewJiraClient(MockPayloadTransporter())

	reply, err := client.ProjectService().Clone(jira.CloneProject{
		Key:                "PROJECT",
		Name:               "PROJECT",
		ExistingProjectKey: "TEMPLATE",
	})
	assert.Nil(t, err)
	assert.Equal(t, "PROJECT", reply.Name)
}

func TestProjectService_Update(t *testing.T) {
	var err error

	var client = NewJiraClient(MockPayloadTransporter())

	reply, err := client.ProjectService().Update("PROJECT", jira.UpdateProject{
		Name: "NEW-PROJECT",
	})
	assert.Nil(t, err)
	assert.Equal(t, "NEW-PROJECT", reply.Name)
}

func TestProjectService_Read(t *testing.T) {
	var err error

	var client = NewJiraClient(MockPayloadTransporter())

	reply, err := client.ProjectService().Read("PROJECT")
	assert.Nil(t, err)
	assert.Equal(t, "PROJECT", reply.Name)
}

func TestProjectService_ReadAll(t *testing.T) {
	var err error

	var client = NewJiraClient(MockPayloadTransporter())

	reply, err := client.ProjectService().ReadAll()
	assert.Nil(t, err)
	assert.Equal(t, "PROJECT", reply[0].Name)
}

func TestProjectService_Delete(t *testing.T) {
	var err error

	var client = NewJiraClient(MockPayloadTransporter())

	success, err := client.ProjectService().Delete("PROJECT", false)
	assert.Nil(t, err)
	assert.True(t, success)
}
