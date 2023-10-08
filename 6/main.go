package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// cпособы остановки горутин:
// 1) завершить main горутину, после этого все другие горутины завершат свою работу:
/*

func main() {
	go func() {

	}()
}

*/
// 2) использовать канал для передачи данных и завершения горутин способ представлен в задаче 4
// или использовать отдельный канал для завершения (к примеру канал stop и вызываем close(stop))
// 3) использовать контекст для завершения горутин(представлен в данном коде):

func main() {
	ctx, cancelCtx := context.WithCancel(context.Background()) //контекст с завершением
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for {
			select {
			case <-ctx.Done(): // контекст завершен, завершаем горутину
				wg.Done()
				return

			default:
				fmt.Println("Out")
				time.Sleep(time.Second)
			}
		}
	}()

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Время пришло...")
		cancelCtx() //завершаем контекст
		wg.Done()
	}()

	wg.Wait()
}
