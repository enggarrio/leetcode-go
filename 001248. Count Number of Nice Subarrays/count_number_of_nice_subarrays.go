package countnumberofnicesubarrays

// https://leetcode.com/problems/count-number-of-nice-subarrays/description/

func numberOfSubarrays(nums []int, k int) int {
	result := 0
	lenNums := len(nums)
	odds := []int{}
	for index := 0; index < lenNums; index++ {
		if nums[index]%2 == 1 {
			odds = append(odds, index+1)
		}
	}

	if len(odds) < k {
		return result
	}

	min := 0
	max := 0
	for index := 0; index < len(odds); index++ {
		if index+1 >= k {
			var left, right int

			// calculate left side number
			if min == 0 {
				left = odds[min]
			} else {
				left = odds[min] - odds[min-1]
			}

			// calculate right side number
			if max == len(odds)-1 {
				right = lenNums - odds[max] + 1
			} else {
				right = odds[max+1] - odds[max]
			}

			result += left * right
			min++
		}
		max++
	}

	return result
}
