package averagewaitingtime

// https://leetcode.com/problems/average-waiting-time/description/

func averageWaitingTime(customers [][]int) float64 {
	var total float64
	current := 0

	for _, customer := range customers {
		if current <= customer[0] {
			total += float64(customer[1])
			current = customer[0] + customer[1]
		} else {
			total += float64(current - customer[0] + customer[1])
			current += customer[1]
		}
	}

	return (total / float64(len(customers)))
}
