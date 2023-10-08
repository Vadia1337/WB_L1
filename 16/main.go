package main

import "fmt"

func main() {
	arr := []int{2, 10, 11, 4, 5, 1}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, low int, high int) {
	if low < high {
		pi := partion(arr, low, high)

		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partion(arr []int, low int, high int) int {
	pivot := arr[high] // опорным элементом выбрали крайний правый элемент массива
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++

			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
