package topics

func InsertionSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		currentvalue := nums[i]
		var j int
		for j = i - 1; j >= 0 && nums[j] > currentvalue; j-- {
			nums[j+1] = nums[j]
		}
		nums[j+1] = currentvalue
	}
	return nums
}
