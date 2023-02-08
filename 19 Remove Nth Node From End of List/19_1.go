package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	slow, fast := head, head
	count := 0
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		if count >= n {
			slow = slow.Next
		}
		count++
	}
	if count+1 != n {
		slow.Next = slow.Next.Next
		return head
	}
	return head.Next
}
