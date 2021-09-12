package main

import "fmt"

func main() {
	str := "gogogo"
	fmt.Printf("src=%s target=%s result=%s\n", str, "str", str)
	fmt.Printf("src=%s target=%s result=%d\n", str, "len", len(str))
	fmt.Printf("src=%s target=%s result=%c\n", str, "index 0", str[0])
	fmt.Printf("src=%s target=%s result=%c\n", str, "index len-1", str[len(str)-1])
	fmt.Printf("src=%s target=%s result=%s\n", str, "hello +", "hello "+str)
}