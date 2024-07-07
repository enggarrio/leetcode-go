package waterbottles

// https://leetcode.com/problems/water-bottles/description/

func numWaterBottles(numBottles int, numExchange int) int {
	result := numBottles

	for numBottles >= numExchange {
		emptyBottles := numBottles % numExchange
		result += (numBottles / numExchange)
		numBottles = numBottles/numExchange + emptyBottles
	}

	return result
}
