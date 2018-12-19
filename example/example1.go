// test project main.go
package main //命令源码文件必须在这里声明自己属于main包
https://studygolang.com/pkgdoc GODOC官网
import (
	"fmt"
	"runtime"
)

func init() {
	fmt.Printf("Map: %v\n", m)
	//通过调用runtime包的代码获取当前机器所运行的操作系统及计算机框架
	//而后通过fmt包的Sprintf方法进行字符串格式化并赋值给变量info
	info = fmt.Sprintf("OS:%s,Arch:%s GOROOT:%s", runtime.GOOS, runtime.GOARCH,runtime.GOROOT())
}

var m map[int]string = map[int]string{1: "A", 2: "B", 3: "C"}

//非局部变量,map类型,已被显示赋值

var info string //非局部变量,string类型,未被显式赋值

func main() {
	fmt.Println(info)
}
