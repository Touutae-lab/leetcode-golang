package main

type ListNode struct {
	val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := head

	totalNodes := 0

	for {
		if head == nil {
			break
		}
		totalNodes += 1
		head = head.Next

	}
	reachNode := (totalNodes / 2) + 1

	for i := 1; i < reachNode; i++ {
		dummy = dummy.Next
	}
	return dummy

}
