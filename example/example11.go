// test project main.go
package main

import (
	"fmt"
)

func PrintNumbers() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\n", i)
	}
}

func main() {

	//demp
	PrintNumbers()
}
