package main

import (
	"fmt"
	"reflect"
)

type TestInt int

func main() {
	var i TestInt
	var j int

	i = TestInt(4)
	j = 4

	fmt.Println(reflect.ValueOf(i).Kind())
	fmt.Println(reflect.ValueOf(j).Kind())
}
