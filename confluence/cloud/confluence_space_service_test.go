package cloud

import (
	"github.com/stretchr/testify/assert"
	"github.com/yunarta/terraform-api-transport/transport"
	"github.com/yunarta/terraform-atlassian-api-client/confluence"
	"testing"
)

func TestSpaceService_Create(t *testing.T) {
	transporter := transport.NewHttpPayloadTransport("https://mobilesolutionworks.atlassian.net", transport.BasicAuthentication{
		Username: "yunarta.kartawahyudi@gmail.com",
		Password: "ATATT3xFfGF0SEMZU9JbDe1_FaGptP2Tg88VJl7eMRT6-zwDX8f30QjKQp_cpg5fakOzTMX96W7UYmfR5bsxsAwmsaItdiiTIHlylRWj7KSZkEEhMiwtE1CdihvYvJ6ieajYY9uIqvM0MRjFiW_h7n1qeoM7B3nv6nXR9rjZYMR9orPRmAEO2zQ=E353F1AA",
	})
	client := NewConfluenceClient(transporter)
	space, err := client.SpaceService().Create(confluence.CreateSpaceRequest{
		Key:  "SK",
		Name: "Space Key",
	})
	assert.Nil(t, err)
	assert.Equal(t, "SK", space.Key)
	assert.Equal(t, "Space Key", space.Name)
}

func TestSpaceService_Delete(t *testing.T) {
	transporter := transport.NewHttpPayloadTransport("https://mobilesolutionworks.atlassian.net", transport.BasicAuthentication{
		Username: "yunarta.kartawahyudi@gmail.com",
		Password: "ATATT3xFfGF0SEMZU9JbDe1_FaGptP2Tg88VJl7eMRT6-zwDX8f30QjKQp_cpg5fakOzTMX96W7UYmfR5bsxsAwmsaItdiiTIHlylRWj7KSZkEEhMiwtE1CdihvYvJ6ieajYY9uIqvM0MRjFiW_h7n1qeoM7B3nv6nXR9rjZYMR9orPRmAEO2zQ=E353F1AA",
	})
	client := NewConfluenceClient(transporter)
	err := client.SpaceService().Delete("SK")
	assert.Nil(t, err)
}
