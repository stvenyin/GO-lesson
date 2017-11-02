// test project main.go
package main

import (
	"fmt"
)

func main() {
	a := 10
	f := func() int {
		a = a * 2
		return 5
	}
	x := []int{10, 5}
	fmt.Println(f())
	
}