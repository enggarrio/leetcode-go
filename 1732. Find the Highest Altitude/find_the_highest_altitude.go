package findthehighestaltitude

// https://leetcode.com/problems/find-the-highest-altitude/description/

func largestAltitude(gain []int) int {
	maxAltitude := 0
	currentAltitude := 0
	for index := 0; index < len(gain); index++ {
		if index == 0 {
			currentAltitude = gain[index]
		} else {
			currentAltitude += gain[index]
		}

		if currentAltitude > maxAltitude {
			maxAltitude = currentAltitude
		}
	}

	return maxAltitude
}
