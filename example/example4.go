// test project main.go
package main

import (
	"fmt"
)

func main() {
	
	
	//
	//slice 切片
	//
	s :=[] int {1,2,3 } 
	//直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3
	//s := arr[:] 
	//初始化切片s,是数组arr的引用
	//s := arr[startIndex:endIndex] 
	//将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片
	//s := arr[startIndex:] 
	//缺省endIndex时将表示一直到arr的最后一个元素
	//s := arr[:endIndex] 
	//缺省startIndex时将表示从arr的第一个元素开始
	//s1 := s[startIndex:endIndex] 
	//通过切片s初始化切片s1
	//s :=make([]int,len,cap) 
	//通过内置函数make()初始化切片s,[]int 标识为其元素类型为int的切片

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
