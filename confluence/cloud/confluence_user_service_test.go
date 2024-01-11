package cloud

import (
	"github.com/stretchr/testify/assert"
	"github.com/yunarta/terraform-api-transport/transport"
	"testing"
)

func TestActorService_ReadUser(t *testing.T) {
	transporter := transport.NewHttpPayloadTransport("https://mobilesolutionworks.atlassian.net", transport.BasicAuthentication{
		Username: "yunarta.kartawahyudi@gmail.com",
		Password: "ATATT3xFfGF0SEMZU9JbDe1_FaGptP2Tg88VJl7eMRT6-zwDX8f30QjKQp_cpg5fakOzTMX96W7UYmfR5bsxsAwmsaItdiiTIHlylRWj7KSZkEEhMiwtE1CdihvYvJ6ieajYY9uIqvM0MRjFiW_h7n1qeoM7B3nv6nXR9rjZYMR9orPRmAEO2zQ=E353F1AA",
	})
	client := NewConfluenceClient(transporter)
	user, err := client.ActorService().ReadUser("yunarta.kartawahyudi@gmail.com")
	assert.Nil(t, err)
	assert.NotNil(t, user)
}
