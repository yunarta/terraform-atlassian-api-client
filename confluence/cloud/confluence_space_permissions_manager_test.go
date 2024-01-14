package cloud

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yunarta/terraform-atlassian-api-client/confluence"
	"testing"
)

func TestSpacePermissionsManager_ReadPermissions(t *testing.T) {
	client := NewConfluenceClient(MockPayloadTransporter())

	manager := SpacePermissionsManager{
		spaceKey: "SP",
		client:   client,
	}
	permissions, err := manager.ReadPermissions()
	assert.Nil(t, err)
	assert.NotNil(t, permissions)
}

func TestSpacePermissionsManager_UpdateUserPermission(t *testing.T) {
	client := NewConfluenceClient(MockPayloadTransporter())

	manager := NewSpaceRoleManager(client, "SK")
	permissions, err := manager.ReadPermissions()
	assert.Nil(t, err)
	assert.NotNil(t, permissions)

	err = manager.UpdateUserPermissions("yunarta.kartawahyudi@gmail.com", []string{
		fmt.Sprintf("%s_%s", confluence.OpRead, confluence.TargetSpace),
		fmt.Sprintf("%s_%s", confluence.OpRead, confluence.TargetPage),
		fmt.Sprintf("%s_%s", confluence.OpDelete, confluence.TargetPage),
	})
}

func TestSpacePermissionsManager_UpdateGroupPermission(t *testing.T) {
	client := NewConfluenceClient(MockPayloadTransporter())

	manager := NewSpaceRoleManager(client, "SK")
	permissions, err := manager.ReadPermissions()
	assert.Nil(t, err)
	assert.NotNil(t, permissions)

	err = manager.UpdateGroupPermissions("site-admins", []string{
		fmt.Sprintf("%s_%s", confluence.OpRead, confluence.TargetSpace),
		fmt.Sprintf("%s_%s", confluence.OpRead, confluence.TargetPage),
		fmt.Sprintf("%s_%s", confluence.OpDelete, confluence.TargetPage),
	})
}

func TestSpacePermissionsManager_RemoveUserPermission(t *testing.T) {
	client := NewConfluenceClient(MockPayloadTransporter())

	manager := NewSpaceRoleManager(client, "SK")
	permissions, err := manager.ReadPermissions()
	assert.Nil(t, err)
	assert.NotNil(t, permissions)

	err = manager.UpdateUserPermissions("yunarta.kartawahyudi@gmail.com", []string{})
}
