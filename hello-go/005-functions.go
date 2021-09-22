package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("匿名返回值")
	fmt.Println(mult_val(2, 4))

	fmt.Printf("\n命名返回值\n")
	fmt.Println(mult_val2(2, 4))

	fmt.Printf("\n传递地址\n")
	n := 0
	mult_val3(7, 8, &n)
	fmt.Println(n)

	fmt.Printf("\n传递变长参数\n")
	say_something("hello:", "lisi", "zhaoliu")
	whos := []string{"tom", "jerry"}
	say_something("hello:", whos...)

	fmt.Printf("\ndefer\n")
	selectUsers()

	fmt.Printf("\ndefer 实现 aop\n")
	sendMsg("send hello")

	fmt.Printf("\n函数传递函数，实现方法复用\n")
	tsSF := timeSpent(slowFun)
	x := tsSF(2)
	fmt.Println(x)
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
	fmt.Printf("before arg: %s\n", arg)
	return arg
} 

func after(arg string) {
	fmt.Printf("after arg: %s\n", arg)
} 

func sendMsg(msg string) {
	defer after(before(msg))
	fmt.Println("exec send msg...")
}

//函数传递函数，实现方法复用

//自定义类型
type MyFun func (num int) int
func timeSpent(innerFunc MyFun) MyFun {
	return func(n int) int {
		start := time.Now()
		res := innerFunc(n)
		fmt.Println("time spent: ", time.Since(start).Seconds())
		return res;
	}
}

func slowFun(n int) int {
	time.Sleep(time.Second*1)
	return n * n;
}