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

	errDistanceTooLarge       = errors.New("maxDistance too large")
	errDistanceMustBePositive = errors.New("maxDistance must be a positive number")
)

func findFirstOccurance(haystack, needle []int) ([]int, error) {
	results, err := findAllOccurances(haystack, needle)
	if err != nil {
		return nil, err
	}

	if len(results) < 1 {
		return []int{}, nil
	}

	return results[0], nil
}

func findFirstOccuranceWithMaxDistanceLimit(haystack, needle []int, maxDistance int) ([]int, error) {
	if maxDistance <= 0 {
		return nil, errDistanceMustBePositive
	}

	if maxDistance > len(haystack) {
		return nil, errDistanceTooLarge
	}

	results, err := findAllOccurances(haystack, needle)
	if err != nil {
		return nil, err
	}

	if len(results) < 1 {
		return []int{}, nil
	}

	for _, res := range results {
		if res[len(needle)-1]-res[0] <= maxDistance {
			return res, nil
		}
	}

	return []int{}, nil
}

var findAllOccurances = func(haystack, needle []int) ([][]int, error) {
	if err := validate(haystack, needle); err != nil {
		return nil, err
	}

	results := [][]int{}
	for j := 0; j < len(haystack)-len(needle)+1; j++ {
		var needleIdx int
		result := make([]int, 0, len(needle))
		for i := j; i < len(haystack); i++ {
			h := haystack[i]
			if contains(h, needle[needleIdx]) {
				result = append(result, i)
				needleIdx++
			}

			if needleIdx >= len(needle) {
				results = append(results, result)
				j = result[0]
				break
			}
		}
	}

	return results, nil
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
