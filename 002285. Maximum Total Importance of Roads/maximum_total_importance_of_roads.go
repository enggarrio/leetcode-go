package maximumtotalimportanceofroads

import "sort"

// https://leetcode.com/problems/maximum-total-importance-of-roads/description/

func maximumImportance(n int, roads [][]int) int64 {
	var result int64
	connections := make([]int, n)

	for _, val := range roads {
		connections[val[0]]++
		connections[val[1]]++
	}
	sort.Ints(connections)

	counter := 1
	for _, val := range connections {
		result += int64(val * counter)
		counter++
	}

	return result
}
