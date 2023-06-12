package utils

import "errors"

func BinarySearch(sortedArray []int, start, end, targetValue, iteration int) (int, int, error) {
	iteration++
	if start >= end {
		return 0, iteration, errors.New("not found")
	}
	var middle int = (start + end) / 2
	if sortedArray[middle] == targetValue {
		return middle, iteration, nil
	} else if sortedArray[middle] > targetValue {
		return BinarySearch(sortedArray, start, middle, targetValue, iteration)
	} else {
		return BinarySearch(sortedArray, middle+1, end, targetValue, iteration)
	}
}
