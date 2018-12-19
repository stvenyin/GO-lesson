package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {


	//runtime.GOMAXPROCS(n) 其中n是整数，

	//指定使用多核的话，goroutins都是在一个线程里的，它们之间通过不停的让出时间片轮流运行，达到类似同时运行的效果


	//多线程时切换,当一个goroutine发生阻塞，Go会自动地把与该goroutine处于同一系统线程的其他goroutines转移到另一个系统线程上去，以使这些goroutines不阻塞

	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
	for _, name := range names {
		go func(who string) {
			fmt.Printf("Hello.%s\n", name)
		}(name)
		runtime.Gosched()
	}

	//optput
	//Hello.Eric
	//Hello.Mark
	//Hello.Robert
	//Hello.Jim
	//Hello.Harry
	time.Sleep(5 *time.Second)
}

