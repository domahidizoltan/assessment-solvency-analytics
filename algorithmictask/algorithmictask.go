package solvencyanalytics

import (
	"strconv"
	"strings"
)

func findFirstOccurance(haystack, needle []int) []int {
	var needleIdx int
	results := make([]int, 0, len(needle))
	for i, h := range haystack {
		if contains(h, needle[needleIdx]) {
			results = append(results, i)
			needleIdx++
		}

		if needleIdx >= len(needle) {
			break
		}
	}

	return results
}

func contains(number int, digit int) bool {
	numberStr := strconv.Itoa(number)
	digitStr := strconv.Itoa(digit)

	return strings.Contains(numberStr, digitStr)
}
