package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mx sync.Mutex
	i  int
}

func (c *Counter) Increment(wg *sync.WaitGroup) {
	c.mx.Lock() // можно использовать atomic, это более производительное решение
	c.i++
	c.mx.Unlock()
	wg.Done()
}

func (c *Counter) GetCount() int {
	return c.i
}

func main() {
	c := Counter{}
	numbOfWorkers := 1000
	wg := sync.WaitGroup{}
	wg.Add(numbOfWorkers)

	for i := numbOfWorkers; i > 0; i-- {
		go c.Increment(&wg)
	}

	wg.Wait()

	fmt.Println(c.GetCount())
}
