package sorting

import "fmt"

func FindSmallest(arr []int) int {
	smallest, smallest_index := arr[0], 0

	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest, smallest_index = arr[i], i

		}
	}
	return smallest_index
}

func SelectionSort(arr []int) []int {
	result_arr := []int{}

	for _, _ = range arr {
		smallest := FindSmallest(arr)
		result_arr = append(result_arr, arr[smallest])

		arr = append(arr[:smallest], arr[smallest+1:]...)

	}
	fmt.Printf("result depois %v", result_arr)
	return result_arr
}
