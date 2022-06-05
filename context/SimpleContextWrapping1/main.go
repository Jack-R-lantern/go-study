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

	if v := ctx.Value("s"); v != nil {
		s := v.(int)
		tick = time.Tick(time.Duration(s) * time.Second)
	}

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
	ctx = context.WithValue(ctx, "s", 2)

	go printTick(ctx)

	time.Sleep(5 * time.Second)
	cancel()

	wg.Wait()
}
