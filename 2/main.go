package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	valuesToSq := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	for _, v := range valuesToSq {
		wg.Add(1)
		go worker(v, &wg)
	}
	wg.Wait()
}

func worker(val int, wg *sync.WaitGroup) {
	_, err := fmt.Println(val * val)
	if err != nil {
		log.Fatal("Ошибка вывода в консоль!")
	}
	wg.Done()
}
