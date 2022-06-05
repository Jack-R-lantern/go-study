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

	d := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	go printTick(ctx)

	time.Sleep(5 * time.Second)
	cancel()

	wg.Wait()
}
