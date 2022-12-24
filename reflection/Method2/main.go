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

	rt := reflect.ValueOf(t)

	m := rt.MethodByName("Increment")
	m2 := rt.MethodByName("String")

	m.Call([]reflect.Value{})
	m.Call([]reflect.Value{})
	fmt.Println("call from reflect.TypeOf: ", m2.Call([]reflect.Value{}))
}
