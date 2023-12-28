package _go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre := &ListNode{0, head}
	var len = 0
	curr := pre
	for head != nil {
		len++
		head = head.Next
	}
	for i := 0; i < len-n; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next
	return pre.Next
}
