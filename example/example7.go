// test project main.go
package main

import (
	"fmt"
)

func main() {


	//golang for 循环处理示例
	var cnt int = 0

	number := 1

	for number < 200 {
		number += 2
	}

	for {
		number++
		cnt++
		fmt.Println(number)
		if cnt == 10{
			break
		}

	}

}
