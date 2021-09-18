package main

import "fmt"

func main() {
	map1 := map[string]int{"one": 1, "two": 2, "three": 3}
	map1["four"] = 4
	fmt.Println(map1, len(map1))
	map2 := map[string]int{}
	fmt.Println(map2["key1"])

	fmt.Println("map 中 key 存在或不存在 val 的值都是默认值")
	//map2["key1"] = 0
	fmt.Println(map2["key1"])
	if v, ok := map2["key1"]; ok {
		fmt.Printf("key1 is %d\n", v)
	} else {
		fmt.Println("key1 is not exit")
	}

	fmt.Println("map 遍历")
	for k, v := range map1 {
		fmt.Printf("key=%s val=%d\n", k, v)
	}

	fmt.Println("map 存放方法")
	fm := map[int]func(num int) int{}
	fm[1] = func(num int) int { return num }
	fm[2] = func(num int) int { return num * num }
	fm[3] = func(num int) int { return num * num * num }
	fmt.Println(fm[1](2))
	fmt.Println(fm[2](2))
	fmt.Println(fm[3](2))

	fmt.Println("map for set")
	mySetTest()
}

func mySetTest() {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 2
	fmt.Println("mySet:",mySet, "key:",n," isExit:",mySet[n])
	mySet[2] = true
	fmt.Println(mySet)
	delete(mySet, 1)
	fmt.Println(mySet)
}
