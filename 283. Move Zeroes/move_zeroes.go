package movezeroes

func moveZeroes(nums []int) {
	lastNonZeroIndex := 0
	for index := 0; index < len(nums); index++ {
		if nums[index] != 0 {
			nums[index], nums[lastNonZeroIndex] = nums[lastNonZeroIndex], nums[index]
			lastNonZeroIndex++
		}
	}

	return
}
