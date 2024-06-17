package removingstarsfromastring

// https://leetcode.com/problems/removing-stars-from-a-string/description/

func removeStars(s string) string {
	var resultStack []byte
	for index := 0; index < len(s); index++ {
		if s[index] == '*' {
			resultStack = resultStack[:len(resultStack)-1]
		} else {
			resultStack = append(resultStack, s[index])
		}
	}

	return string(resultStack)
}
