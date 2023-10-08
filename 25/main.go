package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Начинаем ждать...")
	fmt.Println("Подождали", Sleep(2*time.Second))
}

func Sleep(d time.Duration) time.Time {
	return <-time.After(d) // ожидает переданное время а затем возвращает текущее время и солько ждали
}
