package problems

func IsSame(nums1, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}

	newMap := make(map[int]bool)
	for _, val := range nums2 {
		newMap[val] = true
	}

	for _, val := range nums1 {
		_, ok := newMap[val*val]
		if !ok {
			return false
		}
	}
	return true
}
