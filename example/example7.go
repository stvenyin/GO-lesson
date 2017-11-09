// test project main.go
package main

import (
	"fmt"
)

func main() {

	number := 1

	for number < 200 {
		number += 2
	}

	for {
		number++
		fmt.Println(number)
	}

}