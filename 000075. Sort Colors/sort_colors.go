package sortcolors

func sortColors(nums []int) {
	var redCount, whiteCount, blueCount int
	for index := 0; index < len(nums); index++ {
		switch nums[index] {
		case 0:
			redCount++
		case 1:
			whiteCount++
		case 2:
			blueCount++
		}
	}

	for index := 0; index < redCount; index++ {
		nums[index] = 0
	}
	for index := redCount; index < redCount+whiteCount; index++ {
		nums[index] = 1
	}
	for index := redCount + whiteCount; index < redCount+whiteCount+blueCount; index++ {
		nums[index] = 2
	}
}
