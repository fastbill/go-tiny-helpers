package contains_test

import (
	"fmt"

	"github.com/fastbill/tiny-helpers/contains"
)

func ExampleString() {
	input := []string{"a", "b", "c"}
	fmt.Println(contains.String(input, "b"))
	fmt.Println(contains.String(input, "z"))

	// Output:
	// true
	// false
}

func ExampleInt() {
	input := []int{-1, 1, 2, 3}
	fmt.Println(contains.Int(input, 2))
	fmt.Println(contains.Int(input, 100))

	// Output:
	// true
	// false
}

func ExampleInt64() {
	input := []int64{-1, 1, 2, 3}
	fmt.Println(contains.Int64(input, 2))
	fmt.Println(contains.Int64(input, 100))

	// Output:
	// true
	// false
}

func ExampleUint64() {
	input := []uint64{1, 2, 3}
	fmt.Println(contains.Uint64(input, 2))
	fmt.Println(contains.Uint64(input, 100))

	// Output:
	// true
	// false
}
