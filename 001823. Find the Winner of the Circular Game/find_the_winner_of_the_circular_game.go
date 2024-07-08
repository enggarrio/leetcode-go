package findthewinnerofthecirculargame

// https://leetcode.com/problems/find-the-winner-of-the-circular-game/description/

func findTheWinner(n int, k int) int {
	if n == 1 {
		return 1
	} else {
		return ((findTheWinner(n-1, k) + k - 1) % n) + 1 // Josephus problem
	}
}
