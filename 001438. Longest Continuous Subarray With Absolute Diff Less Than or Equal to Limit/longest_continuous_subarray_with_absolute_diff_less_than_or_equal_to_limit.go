package longestcontinuoussubarraywithabsolutedifflessthanorequaltolimit

// https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/description/

func getMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func getAbs(num int) int {
	if num < 0 {
		num *= -1
	}
	return num
}

func longestSubarray(nums []int, limit int) int {
	lenNums := len(nums)
	var left, right, minNumIndex, maxNumIndex int
	result := 1

	for right < lenNums {
		if left == right {
			right++
		} else {
			if nums[right] >= nums[maxNumIndex] {
				maxNumIndex = right
			}

			if nums[right] <= nums[minNumIndex] {
				minNumIndex = right
			}

			if nums[maxNumIndex]-nums[minNumIndex] <= limit {
				right++
			} else {
				result = getMax(result, right-left)
				left = getMin(maxNumIndex, minNumIndex) + 1
				right = left + 1
				minNumIndex = left
				maxNumIndex = left
			}
		}
	}

	result = getMax(result, right-left)
	return result
}
