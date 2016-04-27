package requesttest

import (
	"testing"

	"github.com/kspichale/golang-samples/request"

	"github.com/stretchr/testify/assert"
)

func TestCaseWithVal(t *testing.T) {
	request := request.New().Uri("uri").Method(request.GET).Build()
	assert.Equal(t, "result", request.Execute())
}
