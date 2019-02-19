package ptr_test

import (
	"fmt"

	"github.com/fastbill/go-tiny-helpers/ptr"
)

func Example() {
	type Product struct {
		ID     *uint64
		Name   *string
		Price  *float64
		Amount *int
	}

	testProduct := Product{
		ID:     ptr.Uint64(123),
		Name:   ptr.String("some test product"),
		Price:  ptr.Float64(12.45),
		Amount: ptr.Int(4),
	}

	fmt.Println(testProduct)
}
