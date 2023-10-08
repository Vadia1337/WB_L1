package main

import (
	"fmt"
	"reflect"
)

func main() {
	recognizeType("ss")
	recognizeType(10)
	recognizeType(true)
	recognizeType(make(chan string))
	recognizeType([]int{})
}

func recognizeType(unknownType interface{}) {
	switch reflect.TypeOf(unknownType).Kind() {
	case reflect.Int:
		fmt.Println("Это число")
	case reflect.String:
		fmt.Println("Это строка")
	case reflect.Bool:
		fmt.Println("Булево значение")
	case reflect.Chan:
		fmt.Println("Это канал")
	default:
		fmt.Println("Не знаю что это ...")
	}
}
