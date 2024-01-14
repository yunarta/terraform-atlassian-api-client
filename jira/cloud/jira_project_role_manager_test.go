package cloud

import (
	"github.com/stretchr/testify/assert"
	"github.com/yunarta/terraform-atlassian-api-client/jira"
	"testing"
)

func TestProjectRoleManager_ReadPermissions(t *testing.T) {
	client := NewJiraClient(MockPayloadTransporter())

	manager := NewProjectRoleManager(client, "P")
	permissions, err := manager.ReadRoles([]string{"Administrators", "Developer", "Contributor"})
	assert.Nil(t, err)
	assert.NotNil(t, permissions)

	users := make(map[string]jira.UserRoles)
	for _, user := range permissions.Users {
		users[user.AccountId] = user
	}

	assert.Equal(t, "yunarta.kartawahyudi@gmail.com", users["557058:32b276cf-1a9f-45ae-b3f5-f850bc24f1b9"].Name)
}

func TestProjectRoleManager_UpdateUserPermission(t *testing.T) {
	client := NewJiraClient(MockPayloadTransporter())

	manager := NewProjectRoleManager(client, "P")
	permissions, err := manager.ReadRoles([]string{"Administrators", "Developer", "Contributor"})
	assert.Nil(t, err)
	assert.NotNil(t, permissions)

	err = manager.UpdateUserRoles("yunarta.kartawahyudi@gmail.com", []string{
		"Contributor",
	})
	manager.Finalized()
}

func TestProjectRoleManager_UpdateGroupPermission(t *testing.T) {
	client := NewJiraClient(MockPayloadTransporter())

	manager := NewProjectRoleManager(client, "P")
	permissions, err := manager.ReadRoles([]string{"Administrators", "Developer", "Contributor"})
	assert.Nil(t, err)
	assert.NotNil(t, permissions)

	err = manager.UpdateGroupRoles("site-admins", []string{
		"Contributor",
	})
	err = manager.UpdateGroupRoles("jira-admins-mobilesolutionworks", []string{
		"Contributor",
	})
	manager.Finalized()
}

//func TestProjectRoleManager_RemoveUserPermission(t *testing.T) {
//	client := NewJiraClient(MockPayloadTransporter())
//
//	manager := NewSpaceRoleManager(client, "SK")
//	permissions, err := manager.ReadPermissions()
//	assert.Nil(t, err)
//	assert.NotNil(t, permissions)
//
//	err = manager.UpdateUserPermissions("yunarta.kartawahyudi@gmail.com", []string{})
//}
