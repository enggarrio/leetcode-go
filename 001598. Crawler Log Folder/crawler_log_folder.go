package crawlerlogfolder

// https://leetcode.com/problems/crawler-log-folder/description/

func minOperations(logs []string) int {
	result := 0

	for _, val := range logs {
		if val[1] == '.' {
			if result > 0 {
				result--
			}
		} else if val[0] != '.' {
			result++
		}
	}

	return result
}
