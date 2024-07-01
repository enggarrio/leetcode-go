package threeconsecutiveodds

// https://leetcode.com/problems/three-consecutive-odds/description/

func threeConsecutiveOdds(arr []int) bool {
	counter := 0
	for _, val := range arr {
		if val%2 == 1 {
			counter++
		} else {
			counter = 0
		}
		if counter == 3 {
			return true
		}
	}
	return false
}
