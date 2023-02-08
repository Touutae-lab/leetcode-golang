package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	node := head
	count := 0

	for node != nil {
		count++
		node = node.Next
	}

	i := count - n

	if i == 0 {
		head = head.Next
		return head
	}

	node = head

	for (i - 1) > 0 {
		node = node.Next
		i--
	}

	node.Next = node.Next.Next

	return head
}
