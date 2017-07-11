package filters

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBrackets_Filter(t *testing.T) {
	b := brackets{}

	actual := b.Filter("()[]{}")

	assert.Equal(t, "", actual)
}

