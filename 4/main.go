package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var workersCount int
	fmt.Println("Укажите кол-во воркеров")
	_, err := fmt.Scanf("%d\n", &workersCount)
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan int, workersCount)

	for n := workersCount; n > 0; n-- {
		go worker(ch)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGINT)

	for {

		if len(sigChan) > 0 {
			fmt.Println(" Нажали ctrl+c")
			close(ch)
			break
		}

		ch <- 10
		time.Sleep(time.Second)
	}
}

func worker(ch chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}
