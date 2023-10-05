package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	arr := [5]int{2, 4, 6, 8, 10}
	res := make(chan int, 5)

	wgOut := sync.WaitGroup{}
	wgOut.Add(1)
	var sum int
	go func() {
		for v := range res {
			sum += v
		}
		wgOut.Done()
	}()

	wgWorkers := sync.WaitGroup{}
	for _, v := range arr {
		wgWorkers.Add(1)
		go worker(v, res, &wgWorkers)
	}
	wgWorkers.Wait()
	close(res)
	wgOut.Wait()

	fmt.Println(sum)

	fmt.Println(time.Since(start))
}

func worker(val int, res chan int, wg *sync.WaitGroup) {
	res <- val * val
	wg.Done()
}
