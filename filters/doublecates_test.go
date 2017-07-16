package filters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoublecates_Filter(t *testing.T) {
	d := doublecates{}

	actual := d.Filter("a b  c    d      e   ")
	assert.Equal(t, "a b c d e ", actual)
}
