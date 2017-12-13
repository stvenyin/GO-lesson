// test project main.go
package main

import (
	"calc" //自定义第三方包
	"fmt"
)

//GO 多态实现

type Brid struct {
}

type LFly interface {
	Fly()
}

func (b *Brid) Fly() {
	fmt.Println("Fly...")
}

func main() {
	var fly LFly = new(Brid)
	fly.Fly()
	fmt.Println(calc.Add(10, 20))

}