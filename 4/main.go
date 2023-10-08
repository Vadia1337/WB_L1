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
		select {
		case <-sigChan:
			fmt.Println(" Нажали ctrl+c")
			close(ch) // закрываем канал
			return

		default:
			ch <- 10
			time.Sleep(time.Second)
		}
	}
}

func worker(ch chan int) {
	for val := range ch { //если канал закрыт, горутина завершает свою работу
		fmt.Println(val)
	}
}

//представленный способ хорошо читается, выглядит компактно и вполне применим на бою, наряду с другими способами:
//завершение контекстом (слушать ctx.Done())
