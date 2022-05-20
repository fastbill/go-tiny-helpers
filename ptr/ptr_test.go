package ptr_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/fastbill/go-tiny-helpers/v2/ptr"
)

func TestString(t *testing.T) {
	result := ptr.String("test")
	assert.Equal(t, "test", *result)
}

func TestInt(t *testing.T) {
	result := ptr.Int(123)
	assert.Equal(t, int(123), *result)
}

func TestInt64(t *testing.T) {
	result := ptr.Int64(123)
	assert.Equal(t, int64(123), *result)
}

func TestUint64(t *testing.T) {
	result := ptr.Uint64(123)
	assert.Equal(t, uint64(123), *result)
}

func TestFloat64(t *testing.T) {
	result := ptr.Float64(123.456)
	assert.Equal(t, float64(123.456), *result)
}

func TestBool(t *testing.T) {
	result := ptr.Bool(true)
	assert.True(t, *result)
}

func TestTime(t *testing.T) {
	now := time.Now()
	result := ptr.Time(now)
	assert.Equal(t, now, *result)
}
