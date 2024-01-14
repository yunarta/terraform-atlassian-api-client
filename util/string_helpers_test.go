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

func TestCoalesceString(t *testing.T) {
	assert.Equal(t, "SecondString", CoalesceString("", "SecondString", "ThirdString"), "First non empty string")
	assert.Equal(t, "", CoalesceString("", "", ""), "All empty strings")
	assert.Equal(t, "FirstString", CoalesceString("FirstString", "SecondString", "ThirdString"), "All non empty strings")
	assert.Equal(t, "LastString", CoalesceString("", "", "LastString"), "Last non empty string")
	assert.Equal(t, "", CoalesceString(), "No input")
}
