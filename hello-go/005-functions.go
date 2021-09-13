package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(mult_val(2, 4))
	fmt.Println(mult_val2(2, 4))
	n := 0
	mult_val3(7, 8, &n)
	fmt.Println(n)

	say_something("hello:", "lisi", "zhaoliu")
	whos := []string{"tom", "jerry"}
	say_something("hello:", whos...)

	selectUsers()

	sendMsg("send hello")
}

//匿名返回值
func mult_val(i, j int) (int, int, int) {
	return i + j, i * j, i - j
}

//命名返回值
func mult_val2(i, j int) (add, mult, sub int) {
	add = i + j
	mult = i * j
	sub = i - j
	return
}

//传递地址
func mult_val3(x, y int, res *int) {
	*res = x * y
}

//传递变长参数
func say_something(prefix string, who ...string) {
	fmt.Printf("%s %s\n", prefix, who)
}

//defer
func selectUsers() {
	fmt.Println("connect to DB")
	defer fmt.Println("closed connect")
	fmt.Println("select * from users")
}

//defer 实现 aop
func before(arg string) string {
	log.Printf("before arg: %s\n", arg)
	return arg
} 

func after(arg string) {
	log.Printf("after arg: %s\n", arg)
} 

func sendMsg(msg string) {
	defer after(before(msg))
	log.Println("exec send msg...")
}