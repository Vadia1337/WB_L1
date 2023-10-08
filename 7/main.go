package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	workersCount := 100
	m := make(map[int]int)
	mx := sync.Mutex{}

	for i := workersCount; i > 0; i-- {
		i := i
		go func() {
			mx.Lock() // блокируем мапу для других горутин
			m[i] = rand.Int()
			mx.Unlock() // разблокируем
		}()
	}

	time.Sleep(2 * time.Second) // дадим время на исполнение горутин

	fmt.Println(m)
}
