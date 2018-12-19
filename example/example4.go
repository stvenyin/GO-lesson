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
	ss1:= s[:]
	//初始化切片s,是数组arr的引用
	//s := arr[startIndex:endIndex]
	//将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片
	//s := arr[startIndex:]
	//缺省endIndex时将表示一直到arr的最后一个元素
	//s := arr[:endIndex]
	//缺省startIndex时将表示从arr的第一个元素开始,不包括endIndex的最后一个元素
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
	fmt.Println(array)				//数组和切片是独立的运行地址空间
	fmt.Println(ss1)

	//
							//完整demo
	//切片声明
	s0 := make([]int, 5, 10)
	fmt.Printf("len(s0)=%d,cap(s0)=%d,yields=%d \n", len(s0), cap(s0), s0)

	arr := [5]int{5, 4, 3, 2, 1}

	//切片可以从数组中复制，
	//array[startIndex:endIndex]表示：获取数组中第startIndex位置到endIndex-1区间的元素. 其切片长度为:endIndex-startIndex
	// startIndex和endIndex为可选项
	//startIndex 不设置时，默认从头开始，endIndex 不设置时，默认截取到末尾

	s1 := arr[:]     //全部，从头到末尾
	s2 := arr[3:]    //从第3+1个到最后
	s3 := arr[:3]    // 从头到 第3+1个(不包括第3+1个)
	s4 := arr[1:3]   //从第1+1个到 第3+1个
	s5 := s1[1:3]    //来自切片s1
	s6 := new([]int) //指针
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)
	fmt.Println("s3=", s3)
	fmt.Println("s4=", s4)

	fmt.Println("s5=", s5)
	fmt.Println("s6=", s6, ",cap=", cap(*s6), ",len=", len(*s6))
	//fmt.Println("s6=", s6, ",cap=", cap(s6), ",len=", len(s6))



}
