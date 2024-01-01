package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueryParam(t *testing.T) {
	param := QueryParam("name", "")
	assert.Equal(t, "", param)

	param = QueryParam("name", "value")
	assert.Equal(t, "&name=value", param)
}
