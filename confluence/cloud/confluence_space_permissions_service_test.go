package cloud

import (
	"github.com/stretchr/testify/assert"
	"github.com/yunarta/terraform-atlassian-api-client/confluence"
	"testing"
)

func TestSpacePermissionsService_Read(t *testing.T) {
	client := NewConfluenceClient(MockPayloadTransporter())

	space, err := client.SpaceService().Read("SK3")
	assert.Nil(t, err)
	assert.NotNil(t, space)

	permissions, err := client.SpacePermissionsService().Read(space.Id)
	assert.Nil(t, err)
	assert.NotNil(t, permissions)

	// list of all required permissions
	requiredPermissions := []string{
		"create_comment", "archive_page", "read_space", "administer_space",
		"delete_attachment", "delete_comment", "create_attachment", "delete_space",
		"create_page", "restrict_content_space", "create_blogpost", "export_space",
		"delete_page", "delete_blogpost",
	}

	userPermissions := map[string]bool{}
	for _, permission := range *permissions {
		if permission.Principal.Id == "557058:32b276cf-1a9f-45ae-b3f5-f850bc24f1b9" {
			userPermissions[permission.Operation.GetSlug()] = true
		}
	}

	// make sure each required permission is present
	for _, requiredPermission := range requiredPermissions {
		assert.True(t, userPermissions[requiredPermission])
	}
}

func TestSpacePermissionsService_Create(t *testing.T) {
	client := NewConfluenceClient(MockPayloadTransporter())

	permission, err := client.SpacePermissionsService().Create("SK3", confluence.AddPermission{
		Subject: confluence.Subject{
			Type: "user",
			Id:   "557058:32b276cf-1a9f-45ae-b3f5-f850bc24f1b9",
		},
		Operation: confluence.AddOperation{
			Key:    "create",
			Target: "page",
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, permission)

	assert.Equal(t, "create", permission.Operation.Key)
	assert.Equal(t, "page", permission.Operation.Target)
	assert.Equal(t, "user", permission.Subject.Type)
	assert.Equal(t, "557058:32b276cf-1a9f-45ae-b3f5-f850bc24f1b9", permission.Subject.Id)
}

func TestSpacePermissionsService_Delete(t *testing.T) {
	client := NewConfluenceClient(MockPayloadTransporter())

	err := client.SpacePermissionsService().Delete("SK3", "3670025")
	assert.Nil(t, err)
}
