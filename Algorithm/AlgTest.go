package main

import "fmt"

type ListNode struct {
	val int
	next *ListNode
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

func print(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.val)
		head = head.next
	}
	fmt.Println()
}

func main() {
	headA := &head
	print(headA)
	node := ListNode{val: 7,next: headA.next.next}
	node2 := ListNode{val: 8,next: &node}
	headB := ListNode{val: 9,next: &node2}
	print(&headB)
	meet := getMeetNode(headA, &headB)
	fmt.Println(meet.val)
}

func getMeetNode(headA , headB *ListNode) *ListNode {
	l1 := headA
	l2 := headB
	for l1 != l2 {
		if l1 == nil {
			l1 = headB
		}else {
			l1 = l1.next
		}
		if l2 == nil {
			l2 = headA
		}else {
			l2 = l2.next
		}
	}
	return l2
}
