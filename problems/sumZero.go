package problems

import "fmt"

func SumZero(nums []int) []int {
	left := 0
	right := len(nums) - 1

	for left < right {
		sum := nums[left] + nums[right]
		fmt.Println(sum)
		if sum == 0 {
			return []int{left, right}
		} else if sum > 0 {
			right--
		} else {
			left++
		}
	}
	return nil
}
