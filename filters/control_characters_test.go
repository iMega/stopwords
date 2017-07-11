package filters

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestControlCharacters_Filter(t *testing.T) {
	f := control_characters{}
	actual := f.Filter("test\ntest\r\ntest\ttest")
	assert.Equal(t, "test test  test test", actual)
}

