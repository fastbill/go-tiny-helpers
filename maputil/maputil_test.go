package maputil_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fastbill/go-tiny-helpers/v2/maputil"
)

func TestKeys(t *testing.T) {
	input := map[string]string{
		"a": "b",
		"c": "d",
	}
	assert.ElementsMatch(t, []string{"a", "c"}, maputil.Keys(input))
}

func TestValues(t *testing.T) {
	input := map[string]string{
		"a": "b",
		"c": "d",
	}
	assert.ElementsMatch(t, []string{"b", "d"}, maputil.Values(input))
}
