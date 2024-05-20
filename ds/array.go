package ds

import (
	"errors"
)

func InsertByPos(arr []int, index int, value int) ([]int, error) {
	if index < 0 || index > len(arr) {
		return nil, errors.New("Index out of range")
	}

	new_arr := make([]int, len(arr)+1)
	copy(new_arr[:index], arr[:index])
	new_arr[index] = value
	copy(new_arr[index+1:], arr[index:])

	return new_arr, nil
}
