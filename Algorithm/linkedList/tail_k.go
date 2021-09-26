package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func print(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ",head.val)
		head = head.next
	}
	fmt.Println()
}

var head ListNode

func init() {
	head = ListNode{val: 1}
	node1 := ListNode{val: 2}
	node2 := ListNode{val: 3}
	node3 := ListNode{val: 4}
	node4 := ListNode{val: 5}
	head.next = &node1
	node1.next = &node2
	node2.next = &node3
	node3.next = &node4
}

func main() {
	print(&head)
	tailK := tailK(&head, 5)
	fmt.Println(tailK)
}

func tailK(head *ListNode, k int) *ListNode {
	fast := head
	for k > 0 {
		fast = fast.next
		k--
	}
	if fast == nil {
		return head
	}
	slow := head
	for fast.next != nil {
		fast = fast.next
		slow = slow.next
	}
	return slow.next
}