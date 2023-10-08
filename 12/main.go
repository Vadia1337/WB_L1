package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]struct{}) // по определению множество должно хранить уникальные значения

	for _, v := range arr { // перекладываем из массива в мапу, повторяющиеся значения не будут записываться в мапу
		m[v] = struct{}{}
	}

	fmt.Println(m) // получили множество
}
