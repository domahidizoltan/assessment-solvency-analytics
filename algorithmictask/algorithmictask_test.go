package solvencyanalytics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFirstOccurance(t *testing.T) {
	for _, s := range []struct {
		name                       string
		haystack, needle, expected []int
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
	} {
		t.Run(s.name, func(t *testing.T) {
			actual := findFirstOccurance(s.haystack, s.needle)
			assert.Equal(t, s.expected, actual)
		})
	}
}
