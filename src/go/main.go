package main

import (
	"fmt"
)

func main() {
	n_train := int(1e6) // how training examples to generate
	n_val := int(1e5)   // how val examples to generate
	n_test := int(1e5)  // how test examples to generate

	fmt.Println(n_test, n_train, n_val)
}
