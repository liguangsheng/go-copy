package copy

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStringToString(t *testing.T) {
	var src = "sudo i love you"
	var dst = "i love you"
	assert.NoError(t, Copy(&dst, src))
	assert.Equal(t, "sudo i love you", dst)
}

func TestInt64ToInt64(t *testing.T) {
	var src int64 = 64
	var dst int64 = 0
	assert.NoError(t, Copy(&dst, src))
	assert.Equal(t, int64(64), dst)
}

func TestTimeToInt64(t *testing.T) {
	var src = time.Now()
	var dst int64 = 0
	assert.NoError(t, Copy(&dst, src))
	assert.Equal(t, src.Unix(), dst)
}

func TestInt64ToTime(t *testing.T) {
	var src = time.Now().Unix()
	var dst time.Time
	assert.NoError(t, Copy(&dst, src))
	assert.Equal(t, src, dst.Unix())
}

func TestStructToStruct(t *testing.T) {
	now := time.Now()
	var src = struct {
		Field1 int
		Field2 string
		Field3 int64
		Field4 float64
		Field5 time.Time
	}{
		Field1: 1,
		Field2: "you are a good guy",
		Field3: 3,
		Field4: 3.141592654,
		Field5: now,
	}
	var dst struct {
		Field1 int
		Field2 string
		Field3 int64
		Field4 float64
		Field5 int64
	}
	cpr := NewCopier()
	assert.NoError(t, cpr.Copy(&dst, src))
	assert.Equal(t, int(1), dst.Field1)
	assert.Equal(t, "you are a good guy", dst.Field2)
	assert.Equal(t, int64(3), dst.Field3)
	assert.Equal(t, float64(3.141592654), dst.Field4)
	assert.Equal(t, now.Unix(), dst.Field5)
}
