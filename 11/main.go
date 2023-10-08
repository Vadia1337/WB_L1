package main

import "fmt"

func main() {
	slice1 := []string{"a", "o", "k", "n", "m", "do", "km", "aw", "d"}
	slice2 := []string{"d", "do", "q", "m", "n", "sfe", "sl", "aw", "b", "w"}

	m1 := getMapFromSlice(slice1)
	m2 := getMapFromSlice(slice2)
	out := findIntersection(m1, m2)

	fmt.Println(out)
}

func getMapFromSlice(slice []string) map[string]struct{} {
	outMap := make(map[string]struct{})
	for _, v := range slice {
		outMap[v] = struct{}{} // заполняем ключ, значение - пустой структурой
	}
	return outMap
}

func findIntersection(m1, m2 map[string]struct{}) map[string]struct{} {
	intersectionMap := make(map[string]struct{})
	for key := range m1 {
		_, ok := m2[key]
		if ok {
			intersectionMap[key] = struct{}{}
		}
	}
	return intersectionMap
}
