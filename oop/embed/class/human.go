package class

import "fmt"

type Human struct {
	// 필드 이름을 소문자로 하여 첫 스코프를 같은 패키지 내로 닫는다
	name string
}

// 생성자
// 구조체가 public이라도 그 필드가 닫힌 스코프의 경우
// 패키지 외부에서 {}을 사용한 초기화는 할 수 없다.
func NewHuman(name string) *Human {
	return &Human{name: name}
}

// 캡슐화
// 메소드 이름을 대문자부터 시작함으로써 public 범위로 한다.
func (human Human) GetName() string {
	return human.name
}

// setter는 수신기를 포인터로 하지 않으면 값이 변경되지 않는다
func (human *Human) SetName(name string) {
	human.name = name
}

func (human Human) Check() {
	fmt.Printf("%s is a Human.\n", human.name)
}

type Man struct {
	*Human // Human을 내장
}

func NewMan(name string) *Man {
	return &Man{Human: NewHuman(name)}
}
