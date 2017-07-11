package filters

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPuctuation_Filter(t *testing.T) {
	p := &puctuation{}

	actual := p.Filter("Cras nisi nisl, porttitor eget rhoncus id, tincidunt eget ante. Phasellus a metus semper, feugiat sapien eget, bibendum nibh.")

	assert.Equal(t, "Cras nisi nisl porttitor eget rhoncus id tincidunt eget ante Phasellus a metus semper feugiat sapien eget bibendum nibh", actual)
}

func TestPuctuation2_Filter(t *testing.T) {
	p := &puctuation{}

	actual := p.Filter("Cras nisi nisl, porttitor eget rhoncus id, tincidunt eget ante. Phasellus a metus semper, feugiat sapien eget, bibendum nibh!")

	assert.Equal(t, "Cras nisi nisl porttitor eget rhoncus id tincidunt eget ante Phasellus a metus semper feugiat sapien eget bibendum nibh", actual)
}

func TestPuctuation3_Filter(t *testing.T) {
	p := &puctuation{}

	actual := p.Filter("Cras nisi nisl, porttitor eget rhoncus id, tincidunt eget ante. Phasellus a metus semper, feugiat sapien eget, bibendum nibh?")

	assert.Equal(t, "Cras nisi nisl porttitor eget rhoncus id tincidunt eget ante Phasellus a metus semper feugiat sapien eget bibendum nibh", actual)
}