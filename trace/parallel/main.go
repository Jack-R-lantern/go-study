package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	// trace 처리
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()

	ts := []*Task{
		&Task{"one", 1 * time.Second, false},
		&Task{"two", 2 * time.Second, true},
		&Task{"three", 3 * time.Second, false},
	}

	var eg errgroup.Group

	for _, t := range ts {
		t := t
		eg.Go(func() error {
			return work(t)
		})
	}

	if err := eg.Wait(); err != nil {
		// 복수의 병렬 처리에서 에러가 나올 경우는 첫번째 에러를 잡는다
		fmt.Println(err.Error())
	}
}
