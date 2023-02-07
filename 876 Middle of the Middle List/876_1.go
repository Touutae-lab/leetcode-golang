package main

type ListNode struct {
	val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	tmp := [100]*ListNode{}
	count := 0
	for head != nil {
		tmp[count] = head
		head = head.Next
		count++
	}
	return tmp[count/2]
}
