package maputil_test

import (
	"fmt"

	"github.com/fastbill/go-tiny-helpers/maputil"
)

func ExampleKeys() {
	input := map[string]string{
		"a": "b",
		"c": "d",
	}
	fmt.Println(maputil.Keys(input))
}

func ExampleValues() {
	input := map[string]string{
		"a": "b",
		"c": "d",
	}
	fmt.Println(maputil.Values(input))
}
