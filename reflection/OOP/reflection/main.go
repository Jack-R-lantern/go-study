package main

import "os"

func main() {
	builder := Builder{}
	Creator(os.Args[1], &builder).Barking()
}
