package minimumnumberofmovestoseateveryone

import "sort"

// https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone/description/

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	var moves, result int
	for index := 0; index < len(seats); index++ {
		moves = students[index] - seats[index]
		if moves < 0 {
			moves *= -1
		}
		result += moves
	}
	return result
}
