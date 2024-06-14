package findpivotindex

// https://leetcode.com/problems/find-pivot-index/description/

func pivotIndex(nums []int) int {
	result := -1
	rightCumulativeSum := 0
	leftCumulativeSum := 0
	arrayLength := len(nums)
	rightCumulativeSumArray := make([]int, arrayLength, arrayLength)
	for index := arrayLength - 1; index >= 0; index-- {
		rightCumulativeSumArray[index] = rightCumulativeSum
		rightCumulativeSum += nums[index]
	}

	for index := 0; index < arrayLength; index++ {
		if leftCumulativeSum == rightCumulativeSumArray[index] {
			result = index
			break
		}
		leftCumulativeSum += nums[index]
	}

	return result
}
