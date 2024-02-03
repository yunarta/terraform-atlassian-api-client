package bitbucket

import (
	"encoding/json"
	"github.com/yunarta/terraform-api-transport/transport"
)

// XBulkJsonPayload struct represents a payload with YAML data.
// It is intended to be used with API requests or responses where the data content is in YAML format.
type XBulkJsonPayload struct {
	Payload any
}

// ContentMust returns the payload data as a byte slice.
// This method is a 'must' version, meaning it assumes the operation will always succeed and does not return an error.
func (m *XBulkJsonPayload) ContentMust() []byte {
	content, _ := m.Content()
	return content
}

// Accept returns the MIME type that this payload expects to receive.
// Here, it is set to 'application/json', indicating that the payload expects to receive JSON-formatted data.
func (m *XBulkJsonPayload) Accept() string {
	return "application/json"
}

// ContentType returns the MIME type of the content this payload contains.
// In this case, it returns 'application/x-yaml', indicating the content type is YAML.
func (m *XBulkJsonPayload) ContentType() string {
	return "application/vnd.atl.bitbucket.bulk+json"
}

// Content returns the payload data as a byte slice and an error.
// In this implementation, it simply converts the YAML string to a byte slice and returns no error.
func (m *XBulkJsonPayload) Content() ([]byte, error) {
	payload, err := json.Marshal(m.Payload)
	return payload, err
}

// This line ensures that XBulkJsonPayload implements the PayloadData interface from the transport package.
// It's a compile-time assertion that validates interface compliance.
var _ transport.PayloadData = &XBulkJsonPayload{}
