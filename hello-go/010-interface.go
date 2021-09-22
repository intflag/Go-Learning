package main

import "fmt"

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {

}

func (gop *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func main() {
	//接口为非入侵性，实现不依赖于接口定义 duck typing
	//所以接口的定义可以包含在接口使用者包内
	var p Programmer
	p = new(GoProgrammer)
	fmt.Println(p.WriteHelloWorld())
}