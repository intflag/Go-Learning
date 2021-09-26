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
	node := reverse(&head)
	print(node)
	dfsNode := dfsReverse(node)
	print(dfsNode)
}

func dfsReverse(head *ListNode) *ListNode{
	if head == nil || head.next == nil {
		return head
	}
	next := head.next
	newHead := dfsReverse(next)
	next.next = head
	head.next = nil
	return newHead
}

func reverse(head *ListNode) *ListNode {
	var node *ListNode
	for head != nil {
		next := head.next
		head.next = node
		node = head
		head = next
	}
	return node
}
