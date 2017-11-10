// test project main.go
package main

import (
	"fmt"
)

func main() {

	ints := []int{1, 2, 3, 4, 5}
	for i, d := range ints {
		fmt.Printf("%d:%d\n", i, d)
	}

	var i, d int
	for i, d = range ints {
		fmt.Printf("%v:%v\n", i, d)
	}
}