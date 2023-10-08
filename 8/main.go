package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	var inputVal int64
	fmt.Println("Введите редактируемое число")
	_, err := fmt.Scanf("%d\n", &inputVal)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(strconv.FormatInt(inputVal, 2))

	var shift uint
	fmt.Println("Какой бит будем редактировать? Нумирация битов с младшего(крайнего правого)")
	_, err = fmt.Scanf("%d\n", &shift)
	if err != nil {
		log.Fatal(err)
	}

	inputVal |= 1 << shift // сдвигаем единицу на введенное число (сдвиг на 3 = 1000) и выполняем операцию или

	fmt.Println("Заменили бит на 1:", strconv.FormatInt(inputVal, 2))

	inputVal &= ^(1 << shift) //сдвигаем единицу на введенное число (сдвиг на 3 = 1000), инвертируем и выполняем операцию и

	fmt.Println("Заменили бит на 0:", strconv.FormatInt(inputVal, 2))
}
