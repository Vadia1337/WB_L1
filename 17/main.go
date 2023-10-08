package main

import "fmt"

func main() {
	arr := []int{-100, -20, -1, 1, 2, 3, 50, 55, 60, 75, 100, 150}
	foundVal := binarySearch(arr, 2) // слайс и искомое число
	fmt.Println(foundVal)            // нащел индекс искомого числа
}

func binarySearch(arr []int, target int) int {
	leftIndex := 0
	rightIndex := len(arr) - 1

	for leftIndex <= rightIndex {
		mid := (leftIndex + rightIndex) / 2 // нашли середину массива
		switch {
		case arr[mid] == target:
			return mid
		case target < arr[mid]:
			rightIndex = mid - 1
		case target > arr[mid]:
			leftIndex = mid + 1
		}
	}
	return -1
}
