package topics

import "fmt"

func Pivot(arr []int, start, end int) []int {
	pivot := arr[start]
	swapIdx := start

	for i := start + 1; i < len(arr); i++ {
		fmt.Println(arr)
		if pivot > arr[i] {
			swapIdx++
			arr[i], arr[swapIdx] = arr[swapIdx], arr[i]
		}
	}

	arr[start], arr[swapIdx] = arr[swapIdx], arr[start]
	return arr
}
