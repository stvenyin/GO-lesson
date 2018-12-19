// test project main.go
package main

import (
	"fmt"
)


//switch 2种表示

func main() {

	num := 123
	switch {
	case num < 100:
		num++
	case num < 200:
		num--
	default:
		num -= 2
	}
	content := "go"
	fmt.Println(num)
	switch content {
	default:
		fmt.Println("Unknown language")
	case "python":
		fmt.Println("An interpreted Lanauage")
	case "go":
		fmt.Println("A compiled language")
	}

}
