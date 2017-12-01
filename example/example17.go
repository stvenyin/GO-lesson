package main

import (
	"fmt"
	_ "time"
)

func test_print(a int) {
	fmt.Print(a)
}

func test_pipe() {
	pipe1 := make(chan int, 1)
	pipe2 := make(chan int, 1)
	pipe1 <- 1
	pipe2 <- 2
	t1 := <-pipe1
	t2 := <-pipe2
	fmt.Println(t1)
	fmt.Println(t2)
}

func main() {

	test_pipe()
	//for i := 0; i < 100; i++ {
	//	go test_print(i)
	//}
	//time.Sleep(time.Second)
}
