package mostprofitassigningwork

// https://leetcode.com/problems/most-profit-assigning-work/description/

import "sort"

type Work struct {
	Difficulty int
	Profit     int
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	works := []Work{}
	maxProfit := 0

	for index := 0; index < len(difficulty); index++ {
		works = append(works, Work{
			Difficulty: difficulty[index],
			Profit:     profit[index],
		})
	}

	sort.Slice(works, func(i, j int) bool {
		return works[i].Profit > works[j].Profit
	})

	sort.Slice(worker, func(i, j int) bool {
		return worker[i] > worker[j]
	})

	worksIndex := 0
	workerIndex := 0
	for workerIndex < len(worker) && worksIndex < len(works) {
		if worker[workerIndex] >= works[worksIndex].Difficulty {
			maxProfit += works[worksIndex].Profit
			workerIndex++
		} else {
			worksIndex++
		}
	}

	return maxProfit
}
