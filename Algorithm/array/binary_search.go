package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5,7,7,8,8,10}
	idx := sort.SearchInts(nums, 6)
	fmt.Println(idx)
}