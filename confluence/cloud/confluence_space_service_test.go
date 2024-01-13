package cloud

import (
	"github.com/stretchr/testify/assert"
	"github.com/yunarta/terraform-atlassian-api-client/confluence"
	"testing"
)

func TestSpaceService_Create(t *testing.T) {
	client := NewConfluenceClient(MockPayloadTransporter())
	space, err := client.SpaceService().Create(confluence.CreateSpace{
		Key:  "SK",
		Name: "Space Key",
	})
	assert.Nil(t, err)
	assert.Equal(t, "SK", space.Key)
	assert.Equal(t, "Space Key", space.Name)
}

func TestSpaceService_Read(t *testing.T) {
	client := NewConfluenceClient(MockPayloadTransporter())
	space, err := client.SpaceService().Read("SK")
	assert.Nil(t, err)
	assert.Equal(t, "SK", space.Key)
	assert.Equal(t, "Space Key", space.Name)
}

func TestSpaceService_Update(t *testing.T) {
	client := NewConfluenceClient(MockPayloadTransporter())
	space, err := client.SpaceService().Update("SK2", confluence.UpdateSpace{
		Name: "Space Key 2",
		Description: confluence.Description{
			Plain: confluence.ContentValue{
				Value: "Description",
			},
		},
	})
	assert.Nil(t, err)
	assert.Equal(t, "SK2", space.Key)

	space, err = client.SpaceService().Read("SK2")
	assert.Nil(t, err)
	assert.Equal(t, "SK2", space.Key)
	assert.Equal(t, "Space Key 2", space.Name)
	assert.Equal(t, "Description", space.Description.Plain.Value)
}

func TestSpaceService_Delete(t *testing.T) {
	client := NewConfluenceClient(MockPayloadTransporter())
	err := client.SpaceService().Delete("SK")
	assert.Nil(t, err)
}
