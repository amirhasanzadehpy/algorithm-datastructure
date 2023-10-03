package main

import (
	"algorithm-datastructure/topics"
	"fmt"
)

func main() {
	//  test sameFreuqncy algo
	//	result := problems.IsSame([]int{1, 2, 3}, []int{1, 4, 9})
	//	result := problems.ValidAnagram("", "")
	//	result := problems.SumZero([]int{-4, -3, -2, -1, 0, 1, 2, 5})
	//	result := problems.CountUniqueValues([]int{1, 2, 3, 4, 4, 4, 7, 7, 12, 12, 13})
	// result := topics.BubbleSort([]int{32, 23, 12, 77, 8})

	//result := topics.BetterBubbleSort([]int{32, 23, 12, 77, 8})

	//	result := topics.SelectionSort([]int{32, 23, 12, 77, 8})
	//	result := topics.InsertionSort([]int{32, 3323, 312, 7, 8})
	//	result := topics.MergeSort([]int{23, 34, 11, 10, 245, 2})
	result := topics.Pivot([]int{23, 0, 11, 10, 245, 2}, 0, 7)
	fmt.Println(result)
}
