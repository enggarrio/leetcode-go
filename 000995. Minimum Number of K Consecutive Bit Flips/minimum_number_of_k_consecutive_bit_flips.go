package minimumnumberofkconsecutivebitflips

// https://leetcode.com/problems/minimum-number-of-k-consecutive-bit-flips/description/

func minKBitFlips(nums []int, k int) int {
	result := 0
	flips := 0
	lenNums := len(nums)

	for index, val := range nums {
		if index >= k && nums[index-k] == 2 {
			flips--
		}

		if flips%2 == val {
			if (index + k) > lenNums {
				result = -1
				break
			}
			nums[index] = 2
			flips++
			result++
		}
	}

	return result
}
