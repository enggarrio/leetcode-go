package deletethemiddlenodeofalinkedlist

// https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/description/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// remove this type before submit answer
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	length := 0
	if head.Next == nil {
		return nil
	} else {
		for node := head; node != nil; node = node.Next {
			length++
		}
		middle := length / 2
		node := head
		for index := 0; index < middle-1; index++ {
			node = node.Next
		}
		node.Next = node.Next.Next
	}
	return head
}
