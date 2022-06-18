package main

import "Embed/class"

func main() {
	man := class.NewMan("bob")
	man.Check() // Human의 메소드를 자신이 구현한 것 처럼 불러낼 수 있다.
	man.SetName("alice")
	man.Check()
}
