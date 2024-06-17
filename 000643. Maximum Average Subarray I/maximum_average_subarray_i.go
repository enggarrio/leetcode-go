package maximumaveragesubarrayi

// https://leetcode.com/problems/maximum-average-subarray-i/description/

func findMaxAverage(nums []int, k int) float64 {
	maxSum := k * -10000
	tempSum := 0
	if len(nums) == k {
		maxSum = 0
		for index := 0; index < len(nums); index++ {
			maxSum += nums[index]
		}
	} else {
		for index := 0; index < len(nums); index++ {
			if index < k {
				tempSum += nums[index]
			} else {
				if index == k {
					maxSum = tempSum
				}
				tempSum = tempSum - nums[index-k] + nums[index]
				if tempSum > maxSum {
					maxSum = tempSum
				}
			}
		}
	}

	result := float64(float64(maxSum) / float64(k))
	return result
}
