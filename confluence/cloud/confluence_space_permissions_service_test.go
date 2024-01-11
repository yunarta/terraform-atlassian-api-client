package cloud

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yunarta/terraform-api-transport/transport"
	"testing"
)

func TestSpacePermissionsService_GetPermissions(t *testing.T) {
	transporter := transport.NewHttpPayloadTransport("https://mobilesolutionworks.atlassian.net", transport.BasicAuthentication{
		Username: "yunarta.kartawahyudi@gmail.com",
		Password: "ATATT3xFfGF0SEMZU9JbDe1_FaGptP2Tg88VJl7eMRT6-zwDX8f30QjKQp_cpg5fakOzTMX96W7UYmfR5bsxsAwmsaItdiiTIHlylRWj7KSZkEEhMiwtE1CdihvYvJ6ieajYY9uIqvM0MRjFiW_h7n1qeoM7B3nv6nXR9rjZYMR9orPRmAEO2zQ=E353F1AA",
	})
	client := NewConfluenceClient(transporter)
	space, err := client.SpaceService().Get("SK")
	assert.Nil(t, err)
	assert.NotNil(t, space)

	permissions, err := client.SpacePermissionsService().GetPermissions(space.Id)
	assert.Nil(t, err)
	assert.NotNil(t, permissions)
	for _, permission := range *permissions {
		if permission.Principal.Id == "5fdc6a7a44065f013ff59b5d" {
			fmt.Printf("Key = %s\n", permission.Operation.Key)
		}
	}
}
