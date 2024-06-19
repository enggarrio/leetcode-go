package findthedifferenceoftwoarrays

// https://leetcode.com/problems/find-the-difference-of-two-arrays/description/

func findDifference(nums1 []int, nums2 []int) [][]int {
	result := [][]int{
		{},
		{},
	}
	map1 := map[int]bool{}
	map2 := map[int]bool{}

	for index := 0; index < len(nums1); index++ {
		map1[nums1[index]] = true
	}
	for index := 0; index < len(nums2); index++ {
		map2[nums2[index]] = true
	}

	for index := 0; index < len(nums1); index++ {
		if _, ok := map2[nums1[index]]; !ok {
			result[0] = append(result[0], nums1[index])
			map2[nums1[index]] = true
		}
	}
	for index := 0; index < len(nums2); index++ {
		if _, ok := map1[nums2[index]]; !ok {
			result[1] = append(result[1], nums2[index])
			map1[nums2[index]] = true
		}
	}

	return result
}
