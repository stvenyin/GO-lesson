// test project main.go
package main

import (
	"errors"
	"fmt"
)

func update(id int, deptment string) bool {
	if id <= 0 {
		return false
	}
	//Ê¡ÂÔÈô¸ÉÌõÓï¾ä
	return true
}

func updateerror(id int, deptment string) error {
	if id <= 0 {
		return errors.New("The id is INVALID!")
	}
	return nil
}

func main() {

	var num int = 1000

	if 100 < num {
		num++
	} else {
		num--
	}
	fmt.Println(update(100, "ho"))
	fmt.Println(updateerror(-1, "ho"))
	fmt.Println(num)

}