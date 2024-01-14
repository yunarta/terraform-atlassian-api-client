package cloud

import (
	"github.com/stretchr/testify/assert"
	"github.com/yunarta/terraform-atlassian-api-client/jira"
	"slices"
	"strings"
	"testing"
)

func TestProjectRoleService_ReadAllRole(t *testing.T) {
	var err error

	var client = NewJiraClient(MockPayloadTransporter())

	roles, err := client.ProjectRoleService().ReadAllRole()
	assert.Nil(t, err)
	assert.Equal(t, "Administrators", roles[0].Name)
	assert.Equal(t, "atlassian-addons-project-access", roles[1].Name)
	assert.Equal(t, "Developer", roles[2].Name)
}

func TestProjectRoleService_ReadProjectRoles(t *testing.T) {
	var err error

	var client = NewJiraClient(MockPayloadTransporter())

	roles, err := client.ProjectRoleService().ReadProjectRoles("P")
	assert.Nil(t, err)

	slices.SortFunc(roles, func(a, b jira.RoleType) int {
		return strings.Compare(a.Name, b.Name)
	})
	assert.Equal(t, "Administrators", roles[0].Name)
	assert.Equal(t, "Contributor", roles[1].Name)
	assert.Equal(t, "Developer", roles[2].Name)
	assert.Equal(t, "atlassian-addons-project-access", roles[3].Name)
}

func TestProjectRoleService_ReadProjectRoleActors(t *testing.T) {
	var err error

	var client = NewJiraClient(MockPayloadTransporter())

	actors, err := client.ProjectRoleService().ReadProjectRoleActors("P", "10008")
	assert.Nil(t, err)

	slices.SortFunc(actors, func(a, b jira.Actor) int {
		return strings.Compare(a.DisplayName, b.DisplayName)
	})

	assert.Equal(t, "Administrator", actors[0].DisplayName)
	assert.Equal(t, "Trello", actors[1].DisplayName)
	assert.Equal(t, "Yunarta Kartawahyudi", actors[2].DisplayName)
}

func TestProjectRoleService_AddProjectRole(t *testing.T) {
	var err error

	var client = NewJiraClient(MockPayloadTransporter())

	err = client.ProjectRoleService().AddProjectRole("P", "10008", []string{"557058:32b276cf-1a9f-45ae-b3f5-f850bc24f1b9"}, []string{})
	assert.Nil(t, err)

	err = client.ProjectRoleService().AddProjectRole("P", "10008", []string{"new-uuid"}, []string{})
	assert.Nil(t, err)
}

func TestProjectRoleService_RemoveProjectRole(t *testing.T) {
	var err error

	var client = NewJiraClient(MockPayloadTransporter())

	err = client.ProjectRoleService().RemoveProjectRole("P", "10008", []string{"557058:32b276cf-1a9f-45ae-b3f5-f850bc24f1b9"}, []string{"557058:32b276cf-1a9f-45ae-b3f5-f850bc24f1b9"})
	assert.Nil(t, err)

	err = client.ProjectRoleService().RemoveProjectRole("P", "10008", []string{"new-uuid"}, []string{})
	assert.Nil(t, err)
}
