package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	var timeSec int
	fmt.Println("Укажите время выполнения программы в секундах")
	_, err := fmt.Scanf("%d\n", &timeSec)
	if err != nil {
		log.Fatal(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	ch := make(chan int)
	timer := time.After(time.Duration(timeSec) * time.Second) // таймер

	go func() { // пишем данные в канал
		for {
			ch <- 1
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() { //читаем из канала
		for {
			select {
			case <-timer: //время закончилось , завершаем горутину , закрываем вейт группу.
				fmt.Println("Время закончилось")
				wg.Done()
				return

			default:
				fmt.Println(<-ch)
			}
		}
	}()

	wg.Wait()
}
