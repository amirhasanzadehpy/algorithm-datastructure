package topics

func SelectionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[min] > nums[j] {
				min = j
			}
		}
		if i != min {
			nums[i], nums[min] = nums[min], nums[i]
		}

	}
	return nums
}
