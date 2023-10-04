package topics

func pivot(arr []int, start, end int) int {
	pivotInt := arr[start]
	swapIdx := start

	for i := start + 1; i < len(arr); i++ {
		if pivotInt > arr[i] {
			swapIdx++
			arr[i], arr[swapIdx] = arr[swapIdx], arr[i]
		}
	}

	arr[start], arr[swapIdx] = arr[swapIdx], arr[start]
	return swapIdx
}

func QuickSort(arr []int, start, end int) {

	if start < end {
		swapIndex := pivot(arr, start, end)

		// left
		QuickSort(arr, start, swapIndex-1)
		// right
		QuickSort(arr, swapIndex+1, end)
	}
}
