package solvencyanalytics

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var errFindAllOccurances = errors.New("findAllOccurances error")

func TestFindAllOccurances(t *testing.T) {
	for _, s := range []struct {
		name             string
		haystack, needle []int
		expected         [][]int
		expectedError    error
	}{
		{
			name:     "all_occurances",
			haystack: []int{662, 154063, 38, 1, 946773, 7877907760054, 332, 76826670, 7653639346039, 90593, 2567954972664},
			needle:   []int{6, 5, 4},
			expected: [][]int{
				{0, 1, 4},
				{1, 5, 8},
				{4, 5, 8},
				{5, 8, 10},
				{7, 8, 10},
				{8, 9, 10},
			},
		},
		{
			name:          "haystack_is_empty",
			haystack:      nil,
			needle:        []int{0},
			expectedError: errHaystackEmpty,
		},
		{
			name:          "needle_is_empty",
			haystack:      []int{0},
			needle:        nil,
			expectedError: errNeedleEmpty,
		},
		{
			name:          "haystack_is_shorter",
			haystack:      []int{123},
			needle:        []int{1, 2},
			expectedError: errHaystackShorter,
		},
	} {
		t.Run(s.name, func(t *testing.T) {
			actual, actualError := findAllOccurances(s.haystack, s.needle)
			assert.Equal(t, s.expected, actual)
			if s.expected != nil {
				assert.ErrorIs(t, s.expectedError, actualError)
			}
		})
	}
}

func TestFindFirstOccurance(t *testing.T) {
	for _, s := range []struct {
		name                       string
		haystack, needle, expected []int
		expectedError              error
		mockedReturn               func([]int, []int) ([][]int, error)
	}{
		{
			name:     "test_1",
			haystack: []int{662, 154063, 38, 1, 946773, 7877907760054, 332, 76826670, 7653639346039, 90593, 2567954972664},
			needle:   []int{6, 5, 4},
			expected: []int{0, 1, 4},
		},
		{
			name:     "test_2",
			haystack: []int{5, 3, 5},
			needle:   []int{3, 5},
			expected: []int{1, 2},
		},
		{
			name:     "no_results",
			haystack: []int{1},
			needle:   []int{2},
			expected: []int{},
		},
		{
			name: "receiving_error",
			mockedReturn: func([]int, []int) ([][]int, error) {
				return nil, errFindAllOccurances
			},
			expectedError: errFindAllOccurances,
		},
	} {
		t.Run(s.name, func(t *testing.T) {
			if s.mockedReturn != nil {
				bkp := findAllOccurances
				defer func() {
					findAllOccurances = bkp
				}()
				findAllOccurances = s.mockedReturn
			}

			actual, actualError := findFirstOccurance(s.haystack, s.needle)
			assert.Equal(t, s.expected, actual)
			if s.expected != nil {
				assert.ErrorIs(t, s.expectedError, actualError)
			}
		})
	}
}

func TestFindFirstOccuranceWithMaxDistanceLimit(t *testing.T) {
	for _, s := range []struct {
		name                       string
		haystack, needle, expected []int
		maxDistance                int
		expectedError              error
		mockedReturn               func([]int, []int) ([][]int, error)
	}{
		{
			name:        "test_1",
			haystack:    []int{662, 154063, 38, 1, 946773, 7877907760054, 332, 76826670, 7653639346039, 90593, 2567954972664},
			needle:      []int{6, 5, 4},
			maxDistance: 3,
			expected:    []int{7, 8, 10},
		},
		{
			name:        "no_results",
			haystack:    []int{1},
			needle:      []int{2},
			maxDistance: 1,
			expected:    []int{},
		},
		{
			name:        "no_results_with_max_distance",
			haystack:    []int{662, 154063, 38, 1, 946773, 7877907760054, 332, 76826670, 7653639346039, 90593, 2567954972664},
			needle:      []int{6, 5, 4},
			maxDistance: 1,
			expected:    []int{},
		},
		{
			name:          "distance_too_large_error",
			maxDistance:   11,
			expectedError: errDistanceTooLarge,
		},
		{
			name:          "distance_positive_error_1",
			maxDistance:   0,
			expectedError: errDistanceMustBePositive,
		},
		{
			name:          "distance_positive_error_2",
			maxDistance:   0,
			expectedError: errDistanceMustBePositive,
		},
		{
			name:        "receiving_error",
			haystack:    []int{0},
			needle:      []int{0},
			maxDistance: 1,
			mockedReturn: func([]int, []int) ([][]int, error) {
				return nil, errFindAllOccurances
			},
			expectedError: errFindAllOccurances,
		},
	} {
		t.Run(s.name, func(t *testing.T) {
			if s.mockedReturn != nil {
				bkp := findAllOccurances
				defer func() {
					findAllOccurances = bkp
				}()
				findAllOccurances = s.mockedReturn
			}

			actual, actualError := findFirstOccuranceWithMaxDistanceLimit(s.haystack, s.needle, s.maxDistance)
			assert.Equal(t, s.expected, actual)
			if s.expected != nil {
				assert.ErrorIs(t, s.expectedError, actualError)
			}
		})
	}
}

func TestFindFirstOccuranceWithMinimumPossibleDistance(t *testing.T) {
	for _, s := range []struct {
		name                       string
		haystack, needle, expected []int
		expectedError              error
		mockedReturn               func([]int, []int) ([][]int, error)
	}{
		{
			name:     "test_1",
			haystack: []int{662, 154063, 38, 1, 946773, 7877907760054, 332, 76826670, 7653639346039, 90593, 2567954972664},
			needle:   []int{6, 5, 4},
			expected: []int{8, 9, 10},
		},
		{
			name:     "no_results",
			haystack: []int{1},
			needle:   []int{2},
			expected: []int{},
		},
		{
			name:     "receiving_error",
			haystack: []int{0},
			needle:   []int{0},
			mockedReturn: func([]int, []int) ([][]int, error) {
				return nil, errFindAllOccurances
			},
			expectedError: errFindAllOccurances,
		},
	} {
		t.Run(s.name, func(t *testing.T) {
			if s.mockedReturn != nil {
				bkp := findAllOccurances
				defer func() {
					findAllOccurances = bkp
				}()
				findAllOccurances = s.mockedReturn
			}

			actual, actualError := findFirstOccuranceWithMinimumPossibleDistance(s.haystack, s.needle)
			assert.Equal(t, s.expected, actual)
			if s.expected != nil {
				assert.ErrorIs(t, s.expectedError, actualError)
			}
		})
	}
}
