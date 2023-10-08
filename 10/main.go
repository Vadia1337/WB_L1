package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := make(map[float64][]float64) // карта, где ключ - группа с шагом 10, значения - набор значений входящих в группу
	for _, v := range arr {
		var group float64
		//if v > 0 {
		//	group = math.Floor(v/10) * 10
		//} else {
		//	group = math.Ceil(v/10) * 10
		//}
		group = math.Trunc(v/10) * 10 // берем целочисленное значение, тем самым можем избавиться от блока условий выше

		m[group] = append(m[group], v) //
	}
	fmt.Println(m)
}
