package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func printTick(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done: ", ctx.Err())
			wg.Done()
			return
		case <-tick:
			fmt.Println("tick")
		}
	}
}

func main() {
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())

	go printTick(ctx)
	time.Sleep(5 * time.Second)

	cancel()
	wg.Wait()
	fmt.Println("main goroutine exit")
}
