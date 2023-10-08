package main

import "fmt"

func main() {
	i := 2 // элемент, который необходимо удалить
	slice := []int{1, 3, 7, 9, 11}
	fmt.Println(delElement1(i, slice))

	slice = []int{1, 3, 7, 9, 11}
	fmt.Println(delElement2(i, slice))
}

func delElement1(index int, slice []int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func delElement2(index int, slice []int) []int {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
