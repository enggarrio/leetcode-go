package grumpybookstoreowner

// https://leetcode.com/problems/grumpy-bookstore-owner/description/

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	result := 0
	max := 0

	for index := 0; index < len(customers); index++ {
		if grumpy[index] == 0 {
			result += customers[index]
		}
	}

	tempAdditional := 0
	for index := 0; index < len(customers); index++ {
		if index >= minutes && grumpy[index-minutes] == 1 {
			tempAdditional -= customers[index-minutes]
		}
		if grumpy[index] == 1 {
			tempAdditional += customers[index]
		}
		if tempAdditional > max {
			max = tempAdditional
		}
	}

	result += max

	return result
}
