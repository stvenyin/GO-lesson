// os project main.go
package main

import (
	"fmt"
	"os"
)

func main() {
	fileinfo, err := os.Stat(`C:\Users\Administrator\Desktop\1.txt`)
	if err != nil {
		panic(err)
	}
	fmt.Println(fileinfo)
	fmt.Println(fileinfo.Name())
	fmt.Println(fileinfo.IsDir())
	fmt.Println(fileinfo.ModTime())
	fmt.Println(fileinfo.Mode())
	fmt.Println(fileinfo.Size())
	fmt.Println(fileinfo.Sys())
}

