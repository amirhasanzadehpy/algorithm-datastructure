package topics

func BubbleSort(nums []int) []int {
	// shrink last item of array
	for i := len(nums); i > 0; i-- {
		for j := 0; j < i-1; j++ {
			// compare
			if nums[j] > nums[j+1] {
				// swap
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums
}

func BetterBubbleSort(nums []int) []int {
	// add flag for knowing is array need for swap or sorted
	var noSwap bool
	// shrink last item of array
	for i := len(nums); i > 0; i-- {
		noSwap = true
		for j := 0; j < i-1; j++ {
			// compare
			if nums[j] > nums[j+1] {
				// swap
				nums[j], nums[j+1] = nums[j+1], nums[j]
				noSwap = false
			}
		}
		if noSwap {
			break
		}
	}

	return nums
}
