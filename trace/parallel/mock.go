package main

import (
	"fmt"
	"time"
)

type Task struct {
	name     string        // task 이름
	duration time.Duration // task 완료까지 필요로하는 시간
	bug      bool          // true의 경우 처리 시에 에러가 발생한다
}

// work 함수는 인수 Task를 실행하고, 이 실행 결과를 돌려준다
func work(t *Task) error {
	time.Sleep(t.duration) // 의사적으로 대기시킨다

	// task에 bug가 포함되어 있는 경우는 에러를 발생시킨다
	if t.bug {
		return fmt.Errorf("err %s", t.name)
	}
	return nil
}
