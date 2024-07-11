package reversesubstringsbetweeneachpairofparentheses

// https://leetcode.com/problems/reverse-substrings-between-each-pair-of-parentheses/description/

func reverseParentheses(s string) string {
	result := ""
	sLength := len(s)
	parenthesesIndexes := []int{}
	parenthesesPairs := make([]int, sLength)
	index := 0
	direction := 1

	for index, val := range s {
		if val == '(' {
			parenthesesIndexes = append(parenthesesIndexes, index)
		} else if val == ')' {
			pairIndex := parenthesesIndexes[len(parenthesesIndexes)-1]
			parenthesesPairs[index] = pairIndex
			parenthesesPairs[pairIndex] = index
			parenthesesIndexes = parenthesesIndexes[:len(parenthesesIndexes)-1]
		}
	}

	for index < sLength {
		if s[index] == '(' || s[index] == ')' {
			direction *= -1
			index = parenthesesPairs[index]
		} else {
			result += string(s[index])
		}
		index += direction
	}

	return result
}
