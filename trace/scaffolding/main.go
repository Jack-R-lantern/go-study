package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"time"
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

	// 메인 처리
	time.Sleep(3 * time.Second) // 3초 대기
	fmt.Fprintf(os.Stdout, "done")
}
