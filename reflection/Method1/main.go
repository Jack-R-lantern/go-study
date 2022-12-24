package main

import (
	"fmt"
	"reflect"
)

type TestStruct struct {
	TestStructValue int
}

func (t TestStruct) String() string {
	return fmt.Sprintf("TestStructValue is %d", t.TestStructValue)
}

func (t *TestStruct) Increment() {
	t.TestStructValue = t.TestStructValue + t.one()
}

func (t *TestStruct) one() int {
	return 1
}

func main() {
	t := &TestStruct{TestStructValue: 1}

	rt := reflect.TypeOf(t)

	numMethod := rt.NumMethod()

	fmt.Println("NumMethod: ", numMethod)

	l := 0
	for l < numMethod {
		m := rt.Method(l)
		fmt.Println("Method1 Name?: ", m.Name)
		fmt.Println("Method1 Type?: ", m.Type)
		fmt.Println("Method1 Func?: ", m.Func)
		fmt.Println("Method1 Index?: ", m.Index)
		fmt.Println("Method1 Pkg?:", m.PkgPath)

		l = l + 1
	}
}
