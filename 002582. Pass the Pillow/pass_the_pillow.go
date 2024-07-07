package passthepillow

// https://leetcode.com/problems/pass-the-pillow/description/

func passThePillow(n int, time int) int {
	remainingTimes := time % (2*n - 2)

	if remainingTimes < n {
		return remainingTimes + 1
	} else {
		return (2*n - 1) - remainingTimes
	}
}
