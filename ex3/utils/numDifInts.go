package utils

import (
	"strconv"
)
func NumDifferentIntegers(word string) int {
	seen := make(map[int]bool)

	for i := 0; i < len(word); i++ {
		if word[i] >= '0' && word[i] <= '9' {
			num := 0
			for ; i < len(word) && word[i] >= '0' && word[i] <= '9'; i++ {
				digit, _ := strconv.Atoi(string(word[i]))
				num = num*10 + digit
			}
			seen[num] = true
		}
	}
	return len(seen)
}
