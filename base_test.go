package progressBar

import (
	"testing"
	"github.com/hzxiao/goutil/assert"
)

func TestBase(t *testing.T) {
	var b = Base{}
	assert.Equal(t, 1, b.PercentageInt(1, 100))
	assert.Equal(t, 99, b.PercentageInt(99, 100))
	assert.NotEqual(t, 20, b.PercentageInt(0, 100))
	assert.Equal(t, 0, b.PercentageInt(100, 0))
}
