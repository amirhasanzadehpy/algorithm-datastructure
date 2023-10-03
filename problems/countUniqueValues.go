package problems

import "fmt"

func CountUniqueValues(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	slower := 0
	faster := 1

	for ; faster < len(nums); faster++ {
		if nums[slower] == nums[faster] {
			faster++
		} else {
			slower++
			nums[slower] = nums[faster]
		}
		fmt.Println(nums)
	}

	return slower + 1

}
