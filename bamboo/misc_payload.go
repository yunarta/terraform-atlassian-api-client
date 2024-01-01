package bamboo

import "github.com/yunarta/terraform-api-transport/transport"

// XYamlPayload struct represents a payload with YAML data.
// It is intended to be used with API requests or responses where the data content is in YAML format.
type XYamlPayload struct {
	Data string
}

// ContentMust returns the payload data as a byte slice.
// This method is a 'must' version, meaning it assumes the operation will always succeed and does not return an error.
func (m *XYamlPayload) ContentMust() []byte {
	return []byte(m.Data)
}

// Accept returns the MIME type that this payload expects to receive.
// Here, it is set to 'application/json', indicating that the payload expects to receive JSON-formatted data.
func (m *XYamlPayload) Accept() string {
	return "application/json"
}

// ContentType returns the MIME type of the content this payload contains.
// In this case, it returns 'application/x-yaml', indicating the content type is YAML.
func (m *XYamlPayload) ContentType() string {
	return "application/x-yaml"
}

// Content returns the payload data as a byte slice and an error.
// In this implementation, it simply converts the YAML string to a byte slice and returns no error.
func (m *XYamlPayload) Content() ([]byte, error) {
	return []byte(m.Data), nil
}

// This line ensures that XYamlPayload implements the PayloadData interface from the transport package.
// It's a compile-time assertion that validates interface compliance.
var _ transport.PayloadData = &XYamlPayload{}
