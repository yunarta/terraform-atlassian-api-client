package cloud

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yunarta/terraform-api-transport/transport"
	"github.com/yunarta/terraform-atlassian-api-client/confluence"
	"testing"
)

func TestSpacePermissionsManager_ReadPermissions(t *testing.T) {
	transporter := transport.NewHttpPayloadTransport("https://mobilesolutionworks.atlassian.net", transport.BasicAuthentication{
		Username: "yunarta.kartawahyudi@gmail.com",
		Password: "ATATT3xFfGF0SEMZU9JbDe1_FaGptP2Tg88VJl7eMRT6-zwDX8f30QjKQp_cpg5fakOzTMX96W7UYmfR5bsxsAwmsaItdiiTIHlylRWj7KSZkEEhMiwtE1CdihvYvJ6ieajYY9uIqvM0MRjFiW_h7n1qeoM7B3nv6nXR9rjZYMR9orPRmAEO2zQ=E353F1AA",
	})
	client := NewConfluenceClient(transporter)
	space, err := client.SpaceService().Get("SK")
	assert.Nil(t, err)
	assert.NotNil(t, space)

	manager := SpacePermissionsManager{
		spaceKey: "SK",
		client:   client,
	}
	permissions, err := manager.ReadPermissions()
	assert.Nil(t, err)
	assert.NotNil(t, permissions)
}
func TestSpacePermissionsManager_UpdateUserPermission(t *testing.T) {
	transporter := transport.NewHttpPayloadTransport("https://mobilesolutionworks.atlassian.net", transport.BasicAuthentication{
		Username: "yunarta.kartawahyudi@gmail.com",
		Password: "ATATT3xFfGF0SEMZU9JbDe1_FaGptP2Tg88VJl7eMRT6-zwDX8f30QjKQp_cpg5fakOzTMX96W7UYmfR5bsxsAwmsaItdiiTIHlylRWj7KSZkEEhMiwtE1CdihvYvJ6ieajYY9uIqvM0MRjFiW_h7n1qeoM7B3nv6nXR9rjZYMR9orPRmAEO2zQ=E353F1AA",
	})
	client := NewConfluenceClient(transporter)
	space, err := client.SpaceService().Get("SK")
	assert.Nil(t, err)
	assert.NotNil(t, space)

	manager := SpacePermissionsManager{
		spaceKey: "SK",
		client:   client,
	}
	permissions, err := manager.ReadPermissions()
	assert.Nil(t, err)
	assert.NotNil(t, permissions)

	err = manager.UpdateUserRoles("yunarta.kartawahyudi@gmail.com", []string{
		fmt.Sprintf("%s_%s", confluence.OperationRead, confluence.TargetSpace),
		fmt.Sprintf("%s_%s", confluence.OperationCreate, confluence.TargetPage),
		fmt.Sprintf("%s_%s", confluence.OperationArchive, confluence.TargetPage),
		fmt.Sprintf("%s_%s", confluence.OperationCreate, confluence.TargetBlogpost),
		fmt.Sprintf("%s_%s", confluence.OperationDelete, confluence.TargetBlogpost),
		fmt.Sprintf("%s_%s", confluence.OperationCreate, confluence.TargetComment),
		fmt.Sprintf("%s_%s", confluence.OperationDelete, confluence.TargetComment),
		fmt.Sprintf("%s_%s", confluence.OperationCreate, confluence.TargetAttachment),
		fmt.Sprintf("%s_%s", confluence.OperationDelete, confluence.TargetAttachment),
		fmt.Sprintf("%s_%s", confluence.OperationExport, confluence.TargetSpace),
		fmt.Sprintf("%s_%s", confluence.OperationAdminister, confluence.TargetSpace),
	})
}
