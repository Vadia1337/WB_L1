package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [5]int{10, 2, 7, 41, 506}

	wg := sync.WaitGroup{}
	wg.Add(1)

	filledChan := fillChanStage(arr)   // этап наполнения канала
	alreadySq := toSqStage(filledChan) // этап возведения в квадрат
	outStage(alreadySq, &wg)           // этап распечатки канала

	wg.Wait()
}

func fillChanStage(arr [5]int) chan int {
	inCh := make(chan int)
	go func() {
		for _, v := range arr {
			inCh <- v
		}
		close(inCh)
	}()
	return inCh
}

func toSqStage(toSq chan int) chan int {
	outCh := make(chan int)
	go func() {
		for v := range toSq {
			outCh <- v * v
		}
		close(outCh)
	}()
	return outCh
}

func outStage(alreadySq chan int, wg *sync.WaitGroup) {
	go func() {
		for v := range alreadySq {
			fmt.Println(v)
		}
		wg.Done()
	}()
}
