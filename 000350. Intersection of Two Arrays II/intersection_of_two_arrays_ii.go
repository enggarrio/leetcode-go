package intersectionoftwoarraysii

// https://leetcode.com/problems/intersection-of-two-arrays-ii/description/

func intersect(nums1 []int, nums2 []int) []int {
	result := []int{}
	mapNumCount := map[int]int{}

	for _, val := range nums1 {
		if _, ok := mapNumCount[val]; !ok {
			mapNumCount[val] = 1
		} else {
			mapNumCount[val]++
		}
	}

	for _, val := range nums2 {
		if _, ok := mapNumCount[val]; ok && mapNumCount[val] > 0 {
			result = append(result, val)
			mapNumCount[val]--
		}
	}

	return result
}
