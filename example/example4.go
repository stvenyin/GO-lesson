// test project main.go
package main

import (
	"fmt"
)

func main() {
	
	
	//
	//slice 切片

	fmt.Println(len([]string{"Go", "Python", "Java", "C", "C++", "PHP"}))
	fmt.Println(cap([]string{"Go", "Python", "Java", "C", "C++", "PHP"}))
	fmt.Println([]string{"Go", "Python", "Java", "C", "C++", "PHP"}[5])

	array := [...]string{"Go", "Python", "java", "C", "C++", "PHP"}
	slice1 := array[:]
	slice1 = append(slice1, "1", "2")
	slice2 := array[3:]
	slice1[0] = "slice1"
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(array)

}
