package requesttest

import (
	"testing"

	"github.com/kspichale/golang-samples/request"

	"github.com/stretchr/testify/assert"
)

func TestCaseWithVal(t *testing.T) {
	request := builder.New().Uri("uri").Method(builder.GET).Build()
	assert.Equal(t, "result", request.Execute())
}
