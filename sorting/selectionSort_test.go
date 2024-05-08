package sorting

import (
	"reflect"
	"testing"
)

func TestFindSmallest(t *testing.T) {
	tests := []struct {
		arr      []int
		expected int
	}{
		{[]int{5, 3, 4, 1, 0}, 4},
		{[]int{14, 9, 3, 128, 4, 100}, 2},
	}

	for _, test := range tests {
		result := FindSmallest(test.arr)

		if result != test.expected {
			t.Errorf("For arr=%v, expected=%d, got=%d", test.arr, test.expected, result)
		}
	}
}

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		arr      []int
		expected []int
	}{
		{[]int{4, 2, 8, 10}, []int{2, 4, 8, 10}},
		{[]int{3, 7, 1, 5, 11}, []int{1, 3, 5, 7, 11}},
	}

	for _, test := range tests {
		copy_arr := make([]int, len(test.arr))
		copy(copy_arr[:], test.arr[:])

		result := SelectionSort(test.arr)

		if !reflect.DeepEqual(test.expected, result) {
			t.Errorf("For arr=%v, expected=%v, got=%v", copy_arr, test.expected, result)
		}
	}
}
