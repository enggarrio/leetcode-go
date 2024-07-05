package findtheminimumandmaximumnumberofnodesbetweencriticalpoints

// https://leetcode.com/problems/find-the-minimum-and-maximum-number-of-nodes-between-critical-points/description/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// delete this before submit
type ListNode struct {
	Val  int
	Next *ListNode
}

func isCriticalPoint(val int, prev int, next int) bool {
	if val < prev && val < next {
		return true
	}
	if val > prev && val > next {
		return true
	}

	return false
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	result := []int{-1, -1}
	prevNode := head
	node := head.Next
	index := 1
	min := 100001
	var cpNums, firstIndex, lastIndex, prevIndex int

	for node != nil {
		if node.Next == nil {
			break
		}
		if isCriticalPoint(node.Val, prevNode.Val, node.Next.Val) {
			cpNums++
			if cpNums == 1 {
				firstIndex = index
				lastIndex = firstIndex
			}
			if cpNums > 1 {
				prevIndex = lastIndex
				lastIndex = index
				if lastIndex-prevIndex < min {
					min = lastIndex - prevIndex
				}
			}
		}
		index++
		prevNode = node
		node = node.Next
	}

	if cpNums < 2 {
		return result
	}
	result[0] = min
	result[1] = lastIndex - firstIndex

	return result
}
