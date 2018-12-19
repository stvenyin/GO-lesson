package main

import "fmt"

type IMessage interface {
		Print()						//void * 指针,和C/C++ 中不存在虚函数指针列表
}
type BaseMessage struct {
	msg string
}

type SubMessage struct {
	BaseMessage
}

func (message *BaseMessage) Print() {

	fmt.Println("baseMessage:msg", message.msg)
}


func (message *SubMessage) Print() {

	fmt.Println("subMessage:msg", message.msg)
}

func interface_use(i IMessage) {

		i.Print()
}

func main() {

	baseMessage := new(BaseMessage)
	baseMessage.msg = "a"
	interface_use(baseMessage)

	SubMessage := new(SubMessage)
	SubMessage.msg = "b"
	interface_use(SubMessage)
}



