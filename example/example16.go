// test project main.go
package main

import (
	"fmt"
	_ "runtime"
)

type Addr struct {
	city     string
	district string
}

type Person struct {
	Name    string
	Age     uint8
	Address Addr
}

func main() {
	p1 := Person{"Harry", 32, Addr{"Beijing", "Haidian"}}
	fmt.Printf("p1 (1):%v\n", p1)
}
