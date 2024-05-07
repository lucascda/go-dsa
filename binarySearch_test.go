package main

import "testing"

func TestBinarySearch(t *testing.T) {

	tests := []struct {
		arr      []int
		target   int
		expected int
	}{

		{[]int{0, 2, 4, 6, 8, 10, 13, 15}, 15, 7},
		{[]int{0, 2, 4, 6, 8, 10, 13, 15}, 0, 0},
		{[]int{0, 3, 5, 7, 11, 17}, 5, 2},
		{[]int{0, 3, 5, 7, 11, 17}, 21, -1},
	}
	for _, test := range tests {
		result, _ := BinarySearch(test.arr, test.target)

		if result != test.expected {
			t.Errorf("For arr=%v, target=%d, expected=%d, got=%d", test.arr, test.target, test.expected, result)
		}
	}
}
