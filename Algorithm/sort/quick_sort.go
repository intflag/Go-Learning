package main

import "fmt"

func main() {
	nums := []int{6, 11, 3, 9, 8}
	fmt.Println(nums)
	quickSort(&nums)
	fmt.Println(nums)
}

func quickSort(nums *[]int) {
	sort(nums, 0, len(*nums)-1)
}

func sort(nums *[]int, left, right int) {
	if left >= right {
		return
	}
	mid := getMid(nums, left, right)
	sort(nums, left, mid-1)
	sort(nums, mid+1, right)
}

func getMid(nums *[]int, left, right int) int {
	mid := left
	flag := (*nums)[right]
	for i := left; i < right; i++ {
		if (*nums)[i] < flag {
			swap(nums, i, mid)
			mid++;
		}
	}
	swap(nums, mid, right)
	return mid
}

func swap(nums *[]int, i,j int) {
	(*nums)[i],(*nums)[j] = (*nums)[j],(*nums)[i]
}
