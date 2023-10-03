package problems

func ValidAnagram(str1, str2 string) bool {
	object := make(map[string]int)
	for _, word := range str1 {
		object[string(word)]++
	}
	for _, word := range str2 {
		object[string(word)]++
	}

	for _, value := range object {
		if value != 2 {
			return false
		}
	}
	return true
}
