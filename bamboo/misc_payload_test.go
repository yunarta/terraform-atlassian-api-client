package bamboo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXYamlPayload_Accept(t *testing.T) {
	xPayload := &XYamlPayload{Data: "test"}

	actual := xPayload.Accept()
	assert.Equal(t, "application/json", actual)
}

func TestXYamlPayload_ContentType(t *testing.T) {
	xPayload := &XYamlPayload{Data: "test"}

	actual := xPayload.ContentType()
	assert.Equal(t, "application/x-yaml", actual)
}

func TestXYamlPayload_Content(t *testing.T) {
	xPayload := &XYamlPayload{Data: "test"}

	actual, err := xPayload.Content()
	assert.Nil(t, err)
	assert.Equal(t, []byte("test"), actual)
}

func TestXYamlPayload_ContentMust(t *testing.T) {
	xPayload := &XYamlPayload{Data: "test"}

	actual := xPayload.ContentMust()
	assert.Equal(t, []byte("test"), actual)
}
