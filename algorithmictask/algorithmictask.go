package solvencyanalytics

import (
	"errors"
	"strconv"
	"strings"
)

var (
	errHaystackEmpty   = errors.New("haystack is empty")
	errNeedleEmpty     = errors.New("needle is empty")
	errHaystackShorter = errors.New("haystack is shorter")
)

func findFirstOccurance(haystack, needle []int) ([]int, error) {
	if err := validate(haystack, needle); err != nil {
		return nil, err
	}

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

	return results, nil
}

func findFirstOccuranceWithMaxDistanceLimit(haystack, needle []int, maximumDistance int) ([]int, error) {
	return nil, nil
}

func validate(haystack, needle []int) error {
	if len(haystack) == 0 {
		return errHaystackEmpty
	}

	if len(needle) == 0 {
		return errNeedleEmpty
	}

	if len(haystack) < len(needle) {
		return errHaystackShorter
	}

	return nil
}

func contains(number int, digit int) bool {
	numberStr := strconv.Itoa(number)
	digitStr := strconv.Itoa(digit)

	return strings.Contains(numberStr, digitStr)
}
