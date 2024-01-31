package solvencyanalytics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAllOccurances(t *testing.T) {
	for _, s := range []struct {
		name             string
		haystack, needle []int
		expected         [][]int
		expectedError    error
	}{
		{
			name:     "test_1",
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
		maximumDistance            int
		expectedError              error
	}{
		{
			name:            "test_1",
			haystack:        []int{662, 154063, 38, 1, 946773, 7877907760054, 332, 76826670, 7653639346039, 90593, 2567954972664},
			needle:          []int{6, 5, 4},
			maximumDistance: 3,
			expected:        []int{7, 8, 10},
		},
	} {
		t.Run(s.name, func(t *testing.T) {
			actual, actualError := findFirstOccuranceWithMaxDistanceLimit(s.haystack, s.needle, s.maximumDistance)
			assert.Equal(t, s.expected, actual)
			if s.expected != nil {
				assert.ErrorIs(t, s.expectedError, actualError)
			}
		})
	}
}
