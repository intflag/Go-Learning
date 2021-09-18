package main

import "fmt"

func main() {
	var arr1 [3]int
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr1, arr2)
	fmt.Println("数组遍历")
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}
	for i, e := range arr2 {
		fmt.Println(i, e)
	}
	for _, e := range arr2 {
		fmt.Println(e)
	}
	fmt.Println("数组截取")
	arr3 := [...]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println("arr3[] ", arr3)
	fmt.Println("arr3[3:] ",arr3[3:])
	fmt.Println("arr3[2:4] ",arr3[2:4])
	fmt.Println("arr3[:3] ",arr3[:3])
	fmt.Println("arr3[:] ",arr3[:])
	
	fmt.Println("数组比较")
	arr4 := [3]int{1,2,3}
	arr5 := [3]int{1,2,3}
	arr6 := [3]int{1,3,3}
	fmt.Println(arr4 == arr5)
	fmt.Println(arr5 == arr6)
}
