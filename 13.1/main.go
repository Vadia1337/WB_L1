package main

import "fmt"

func main() {
	i1 := 1
	i2 := 2

	i1 = i1 + i2
	i2 = i1 - i2
	i1 = i1 - i2

	// тоже самое можно провернуть с умножением

	fmt.Println(i1, i2)
}
