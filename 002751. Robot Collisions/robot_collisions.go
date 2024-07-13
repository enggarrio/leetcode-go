package robotcollisions

// https://leetcode.com/problems/robot-collisions/description/

import "sort"

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	result := []int{}
	numOfRobots := len(positions)
	rightRobots := []int{}
	indexes := make([]int, numOfRobots)
	mapPositionIndex := make(map[int]int, numOfRobots)
	sortedPosition := []int{}

	for index := 0; index < numOfRobots; index++ {
		mapPositionIndex[positions[index]] = index
	}

	sortedPosition = append(sortedPosition, positions...)
	sort.Ints(sortedPosition)

	for index := 0; index < numOfRobots; index++ {
		indexes[index] = mapPositionIndex[sortedPosition[index]]
	}

	for _, index := range indexes {
		if directions[index] == 'R' {
			rightRobots = append(rightRobots, index)
		} else {
			for len(rightRobots) > 0 && healths[index] > 0 {
				rightRobotIndex := rightRobots[len(rightRobots)-1]

				if healths[rightRobotIndex] > healths[index] {
					healths[rightRobotIndex]--
					healths[index] = 0
				} else if healths[rightRobotIndex] < healths[index] {
					healths[rightRobotIndex] = 0
					healths[index]--
					rightRobots = rightRobots[:len(rightRobots)-1]
				} else {
					healths[index] = 0
					healths[rightRobotIndex] = 0
					rightRobots = rightRobots[:len(rightRobots)-1]
				}
			}
		}
	}

	for _, val := range healths {
		if val > 0 {
			result = append(result, val)
		}
	}

	return result
}
