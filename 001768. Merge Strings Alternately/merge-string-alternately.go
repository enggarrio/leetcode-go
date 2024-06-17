package mergestringalternately

func mergeAlternately(word1 string, word2 string) string {
	maxLength := len(word1)
	if len(word2) > maxLength {
		maxLength = len(word2)
	}
	result := ""

	for i := 0; i < maxLength; i++ {
		if i < len(word1) {
			result += string(word1[i])
		}

		if i < len(word2) {
			result += string(word2[i])
		}
	}

	return result
}
