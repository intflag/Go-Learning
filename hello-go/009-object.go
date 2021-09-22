package main

import (
	"fmt"
	"unsafe"
)

//定义 user 对象
type User struct {
	Id   int
	Name string
	age  int
}

//定义 user 的方法，使用指针
func (u *User) String() string {
	return fmt.Sprintf("Type: %T  NameAddress: %x Id:%d Name:%s Age:%d", u, unsafe.Pointer(&u.Name), u.Id, u.Name, u.age)
}

//定义 user 的方法，不使用指针
func (u User) String2() string {
	return fmt.Sprintf("Type: %T  NameAddress: %x Id:%d Name:%s Age:%d", u, unsafe.Pointer(&u.Name), u.Id, u.Name, u.age)
}

func main() {
	fmt.Printf("\n初始化1，相当于java空参构造\n")
	u1 := User{}
	fmt.Printf("tpye: %T\n", u1)
	fmt.Println(u1)
	u1.age = 19
	fmt.Println(u1)

	fmt.Printf("\n初始化2，相当于java全参构造\n")
	u2 := User{1, "tome", 18}
	fmt.Printf("tpye: %T\n", u2)
	fmt.Println(u2)

	fmt.Printf("\n初始化3，相当于java有限参数个数构造\n")
	u3 := User{Name: "jerry", age: 21}
	fmt.Printf("tpye: %T\n", u3)
	fmt.Println(u3)
	u3.Name = "jerry2"
	fmt.Println(u3)

	fmt.Printf("\n初始化4，使用new关键字，返回指针\n")
	u4 := new(User)
	fmt.Printf("tpye: %T\n", u4)
	fmt.Println(u4)

	fmt.Printf("\n对象方法\n")
	fmt.Println("u2: ", u2.String())
	fmt.Println("u3: ", u3.String())

	fmt.Printf("\n方法使用指针与不使用指针的区别\n")
	fmt.Printf("NameAddress: %x\n",unsafe.Pointer(&u3.Name))
	fmt.Println(u3.String())
	fmt.Println(u3.String2())
}
