package main

import "polymorphism/class"

func main() {
	human := class.NewHuman("alice")
	man := class.NewMan("bob")
	woman := class.NewWoman("emily")

	// Human도 Woman도 Speaker 인터페이스로 다룰 수 있다.
	speak(human)
	speak(man)
	speak(woman)
}

func speak(speaker class.Speaker) {
	speaker.Speak()
}
