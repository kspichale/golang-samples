package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaseWithVal(t *testing.T) {
	request := New().Uri("uri").Build()
	assert.Equal(t, "result", request.Execute())
}
