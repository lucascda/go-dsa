package ds

import (
	"reflect"
	"testing"
)

func TestInsertByPos(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		arr      []int
		index    int
		element  int
		expected []int
	}{
		{"InsertOneElement", 5, []int{0, 2, 4, 6, 8}, 2, 10, []int{0, 2, 10, 4, 6, 8}},
		{"InsertOneElement", 5, []int{0, 2, 4, 6, 8}, 3, 7, []int{0, 2, 4, 7, 6, 8}},
		{"InsertOneElement", 5, []int{0, 2, 4, 6, 8}, 4, 5, []int{0, 2, 4, 6, 5, 8}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//arr := make([]int, test.size)
			result_arr, _ := InsertByPos(test.arr, test.index, test.element)

			if !reflect.DeepEqual(test.expected, result_arr) {
				t.Errorf("For arr=%v, expected new arr=%v, but got %v", test.arr, test.expected, result_arr)
			}
		})
	}
}

func TestRemoveByIndex(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		arr      []int
		index    int
		expected []int
	}{
		{"RemoveElementIndex2", 5, []int{0, 2, 4, 6, 8}, 2, []int{0, 2, 6, 8}},
		{"RemoveElementIndex3", 5, []int{0, 2, 4, 6, 8}, 3, []int{0, 2, 4, 8}},
		{"RemoveElementIndex4", 5, []int{0, 2, 4, 6, 8}, 4, []int{0, 2, 4, 6}},
		{"RemoveElementIndex0", 5, []int{0, 2, 4, 6, 8}, 0, []int{2, 4, 6, 8}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result_arr, _ := RemoveByIndex(test.arr, test.index)

			if !reflect.DeepEqual(test.expected, result_arr) {
				t.Errorf("For arr=%v, expected new arr=%v, but got %v", test.arr, test.expected, result_arr)
			}
		})
	}
}
