// test project main.go
package main

import (
	"fmt"
)

func main() {
	
	//数组示例 var []array{}
	
	array := [8]string{1: "Go", "Python", 4: "Java", "C", "C++", "PHP"}
	fmt.Println([...]string{1: "Go", "Python", 4: "Java", "C", "C++", "PHP"}[1])
	fmt.Println(len([8]string{1: "Go", "Python", 4: "Java", "C", "C++", "PHP"}))
	fmt.Println(array[3])
	
	
}
