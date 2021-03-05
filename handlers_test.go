package copy

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimeToInt64(t *testing.T) {
	var src = time.Now()
	var dest int64 = 0
	assert.NoError(t, Copy(&dest, src))
	assert.Equal(t, src.Unix(), dest)
}

func TestInt64ToTime(t *testing.T) {
	var src = time.Now().Unix()
	var dest time.Time
	assert.NoError(t, Copy(&dest, src))
	assert.Equal(t, dest.Unix(), src)
}
