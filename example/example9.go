// test project main.go
package main

import (
	"fmt"
)

func main() {

	n := 100

	if n%3 != 0 {
		goto L1
	}
	switch {
	case n%7 == 0:
		n = 200
		fmt.Printf("%v is a common multiple of 7 and 3.\n", n)
	default:
	}
L1:
	fmt.Printf("%v isn't a multiple of 3.\n", n)

}
