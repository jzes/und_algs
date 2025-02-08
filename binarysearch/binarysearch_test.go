package binarysearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{"Target at the beginning", []int{1, 2, 3, 4, 5}, 1, 0},
		{"Target in the middle", []int{1, 2, 3, 4, 5}, 3, 2},
		{"Target at the end", []int{1, 2, 3, 4, 5}, 5, 4},
		{"Target not in array", []int{1, 2, 3, 4, 5}, 6, -1},
		{"Empty array", []int{}, 1, -1},
		{"Single element - found", []int{10}, 10, 0},
		{"Single element - not found", []int{10}, 5, -1},
		{"Even number of elements", []int{2, 4, 6, 8, 10, 12}, 8, 3},
		{"Odd number of elements", []int{1, 3, 5, 7, 9}, 7, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BinarySearch(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
