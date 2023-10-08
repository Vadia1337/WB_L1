package main

import (
	"fmt"
	"sync"
)

func main() {

	arr := [5]int{2, 4, 6, 8, 10}
	res := 0
	wg := sync.WaitGroup{}
	m := sync.Mutex{}

	for _, v := range arr {
		wg.Add(1)
		go worker(v, &res, &wg, &m)
	}

	wg.Wait()

	fmt.Println(res)
}

func worker(val int, res *int, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	*res += val * val
	m.Unlock()
	wg.Done()
}
