package main

import "fmt"

func main() {
	var s1 []int
	fmt.Println("slice 扩容")
	fmt.Println(len(s1), cap(s1))
	s1 = append(s1, 1)
	fmt.Println(len(s1), cap(s1))
	s1 = append(s1, 2)
	fmt.Println(len(s1), cap(s1))
	s1 = append(s1, 3)
	fmt.Println(len(s1), cap(s1))
	s1 = append(s1, 4)
	fmt.Println(len(s1), cap(s1))
	s1 = append(s1, 5)
	fmt.Println(len(s1), cap(s1))
	
	fmt.Println("slice make")
	s2 := make([]int, 3, 8)
	fmt.Println(len(s2), cap(s2))
	//fmt.Println(s2[4])

	fmt.Println("切片")
	s1 = append(s1, 2,3)
	fmt.Println(s1)

	ps1 := &s1;
	x1 := (*ps1)[2]
	fmt.Println(ps1)
	fmt.Println(x1)
}