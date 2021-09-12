package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("当前时间\n")
	fmt.Println(time.Now())
	
	fmt.Printf("\n时间格式化为字符串\n")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	
	fmt.Printf("\n字符串转时间\n")
	var timeStr1 string = "1996-04-07 00:00:00"
	time1,_ := time.Parse("2006-01-02 15:04:05",timeStr1)
	fmt.Println(time1)
	
	fmt.Printf("\n求时间戳\n")
	fmt.Println(time1.Unix()) //828835200
	
	fmt.Printf("\n时间戳转时间\n")
	fmt.Println(time.Unix(828835200,0))
	
	fmt.Printf("\n计算时间差\n")
	fmt.Println(time.Now().Sub(time1))
}
