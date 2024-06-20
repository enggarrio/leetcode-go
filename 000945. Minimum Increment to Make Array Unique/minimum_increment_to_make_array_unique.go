package minimumincrementtomakearrayunique

// https://leetcode.com/problems/minimum-increment-to-make-array-unique/description/

import "sort"

func minIncrementForUnique(nums []int) int {
	var result int
	sort.Ints(nums)
	for index := 1; index < len(nums); index++ {
		if nums[index] <= nums[index-1] {
			diff := nums[index] - nums[index-1]
			if diff < 0 {
				diff *= -1
			}
			nums[index] += diff + 1
			result += diff + 1
		}
	}

	return result
}
