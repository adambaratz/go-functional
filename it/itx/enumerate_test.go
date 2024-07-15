package itx_test

import (
	"fmt"
	"slices"

	"github.com/BooleanCat/go-functional/v2/it/itx"
)

func ExampleIterator_Enumerate() {
	for index, value := range itx.From(slices.Values([]int{1, 2, 3})).Enumerate() {
		fmt.Println(index, value)
	}

	// Output:
	// 0 1
	// 1 2
	// 2 3
}
