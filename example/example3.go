// test project main.go
package main

import (
	"fmt"
)

func main() {

	//数组示例 var []array{}

	array := [8]string{1: "Go", "Python", 4: "Java", "C", "C++", "PHP"}
	fmt.Println([...]string{1: "Go", "Python", 4: "Java", "C", "C++", "PHP"}[1])
	fmt.Println([...]string{1: "Go", "Python", 4: "Java", "C", "C++", "PHP"}[2])
	fmt.Println([...]string{1: "Go", "Python", 4: "Java", "C", "C++", "PHP"}[3])
	fmt.Println([...]string{1: "Go", "Python", 4: "Java", "C", "C++", "PHP"}[5])
	fmt.Println(len([900]string{1: "Go", "Python", 4: "Java", "C", "C++", "PHP"})) //初始化方式略有不同
	fmt.Println(cap([900]string{1: "Go", "Python", 4: "Java", "C", "C++", "PHP"}))
	fmt.Println(array[3])

}
