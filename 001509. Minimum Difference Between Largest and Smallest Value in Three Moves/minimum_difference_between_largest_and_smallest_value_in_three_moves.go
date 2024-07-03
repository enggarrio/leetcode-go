package minimumdifferencebetweenlargestandsmallestvalueinthreemoves

// https://leetcode.com/problems/minimum-difference-between-largest-and-smallest-value-in-three-moves/description/

import "sort"

func minDifference(nums []int) int {
	var result int
	lenNums := len(nums)
	if lenNums <= 4 {
		return 0
	}

	sort.Ints(nums)
	left := 1
	right := lenNums - 3
	result = nums[right-1] - nums[left-1]
	for left < 4 {
		if nums[right]-nums[left] < result {
			result = nums[right] - nums[left]
		}
		left++
		right++
	}

	return result
}
