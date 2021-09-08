package contains_test

import (
	"fmt"

	"github.com/fastbill/go-tiny-helpers/v2/contains"
)

func ExampleString() {
	input := []string{"a", "b", "c"}
	fmt.Println(contains.String(input, "b"))
	fmt.Println(contains.String(input, "C"))

	// Output:
	// true
	// false
}

func ExampleStringWithEqualFold() {
	input := []string{"APPLE", "banana", "CasheW"}
	fmt.Println(contains.StringWithEqualFold(input, "apple"))
	fmt.Println(contains.StringWithEqualFold(input, "BAnana"))
	fmt.Println(contains.StringWithEqualFold(input, "cashew"))
	fmt.Println(contains.StringWithEqualFold(input, "app"))

	// Output:
	// true
	// true
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
