package ptr_test

import (
	"testing"

	"github.com/fastbill/tiny-helpers/ptr"
	"github.com/stretchr/testify/assert"
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