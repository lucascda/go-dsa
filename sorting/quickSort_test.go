package sorting

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected []int
	}{
		{"QuickSort_1", []int{1, 1, 1, 0, 5, 4}, []int{0, 1, 1, 1, 4, 5}},
		{"QuickSort_2", []int{}, []int{}},
		{"QuickSort_3", []int{3, 1}, []int{1, 3}},
		{"QuickSort_4", []int{7, 3, 1, 0}, []int{0, 1, 3, 7}},
		{"QuickSort_1", []int{-7, -5, 1, -5, 0, 5, 4}, []int{-7, -5, -5, 0, 1, 4, 5}},
	}

	for _, test := range tests {
		arr := []int{}
		arr = append(arr, test.arr...)
		result := QuickSort(test.arr)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For arr=%v, got=%v, expected=%v", arr, result, test.expected)
		}
	}
}
