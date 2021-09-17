package main

import "fmt"

func main() {
	arr := []int{5, 11, 7, 9, 8}
	fmt.Println(arr)
	selectSort(arr)
	fmt.Println(arr)
}

func selectSort(nums []int) {
	for i:=0; i < len(nums); i++ {
		for j:=i; j < len(nums); j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}