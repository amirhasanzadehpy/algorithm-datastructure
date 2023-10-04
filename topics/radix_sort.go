package topics

import (
	"fmt"
	"math"
)

func getDigit(number, i int) int {
	result := (number / int(math.Pow(10, float64(i)))) % 10
	return result
}

func digitCount(number int) int {
	if number == 0 {
		return 1
	}

	return int(math.Log10(float64(number))) + 1
}

func mostDigits(number []int) int {
	maxDigit := 0
	for i := 0; i < len(number); i++ {
		if maxDigit < digitCount(number[i]) {
			maxDigit = digitCount(number[i])
		}
	}

	return maxDigit
}

func RadixSort(arr []int) []int {
	// find biggest number length
	maxDigitcount := mostDigits(arr)
	// loop over biggest length number
	for k := 0; k < maxDigitcount; k++ {
		// make 0-1-2-3-4-5-6-7-8-9 buckets in every loop
		digitBucket := [][]int{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}}
		// loop over items in origin array to put them in bucket
		for i := 0; i < len(arr); i++ {
			// find kth number of item with getDigit function
			// ex : loop 2: item = 123 -> digit is 1
			digit := getDigit(arr[i], k)
			// put the item in specefic bucket about kth number
			digitBucket[digit] = append(digitBucket[digit], arr[i])
		}
		var nums []int
		for _, bucket := range digitBucket {
			nums = append(nums, bucket...)
		}

		arr = nums
		fmt.Println(arr)
	}

	return arr
}
