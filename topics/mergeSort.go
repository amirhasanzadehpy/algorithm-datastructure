package topics

func merge(nums1, nums2 []int) []int {
	result := []int{}
	i := 0
	j := 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			result = append(result, nums2[j])
			j++
		} else {
			result = append(result, nums1[i])
			i++
		}
	}

	for i < len(nums1) {
		result = append(result, nums1[i])
		i++
	}

	for j < len(nums2) {
		result = append(result, nums2[j])
		j++
	}
	return result
}

func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := MergeSort(nums[0:mid])
	right := MergeSort(nums[mid:])

	return merge(left, right)
}
