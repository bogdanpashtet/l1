// Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
// которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker(channel chan int, number int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("\nWorker %v stopped.", number)
			return
		case data := <-channel:
			fmt.Printf("Worker %v got: %v\n", number, data)
		}
	}
}

func main() {
	var ch = make(chan int)
	var count int
	fmt.Print("Enter count of workers: ")
	fmt.Fscan(os.Stdin, &count)

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	i := 0

	for i := 1; i <= count; i++ {
		go worker(ch, i, ctx)
	}

	for {
		select {
		case <-ctx.Done():
			close(ch)
			os.Exit(100)
		default:
			ch <- i
			i++
			time.Sleep(500 * time.Millisecond)
		}
	}
}
