package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func printParent(pctx context.Context) {
	cctx, cancel := context.WithCancel(pctx)
	tick := time.Tick(time.Second)
	var cnt int = 0
	go printChild(cctx)
	for {
		select {
		case <-pctx.Done():
			fmt.Println("Parent Done: ", pctx.Err())
			wg.Done()
			return
		case <-tick:
			fmt.Println("Parent Tick")
			cnt++
			if cnt == 3 {
				cancel()
			}
		}
	}
}

func printChild(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Child Done: ", ctx.Err())
			wg.Done()
			return
		case <-tick:
			fmt.Println("Child Tick")
		}
	}
}

func main() {
	wg.Add(2)
	ctx, cancel := context.WithCancel(context.Background())

	go printParent(ctx)
	time.Sleep(10 * time.Second)

	cancel()
	wg.Wait()
}
