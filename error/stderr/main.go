package main

import (
	"fmt"
	"github.com/pkg/errors"
)

var (
	innerError = &innerErr{Msg: "innerMsg"}
	myError    = &myErr{Msg: "myMsg"}
)

// Custrom errors
type innerErr struct {
	Msg string
}

func (i *innerErr) Error() string {
	return "innerError"
}

type myErr struct {
	Msg string
}

func (m *myErr) Error() string {
	return "myError"
}

// Functions
func innerFunc() error {
	return innerError
}

func middleFunc() error {
	if err := innerFunc(); err != nil {
		return errors.Wrap(err, "middleError")
	}
	return nil
}

func outerFunc() error {
	if err := middleFunc(); err != nil {
		return errors.Wrap(err, "outerError")
	}
	return nil
}

func main() {
	// Get a wrapped error
	outerErr := outerFunc()

	// Cause
	fmt.Printf("\n--- Cause ---\n")
	fmt.Printf("%+v\n", errors.Cause(outerErr))

	// Stack
	fmt.Printf("\n--- Stack ---\n")
	fmt.Printf("%+v\n", outerErr)

	// Is (Compare)
	fmt.Printf("\n--- Is ---\n")
	if errors.Is(outerErr, innerError) {
		fmt.Printf("innerError true\n") // Print
	} else {
		fmt.Printf("innerError false\n")
	}
	if errors.Is(outerErr, myError) {
		fmt.Printf("myError true\n")
	} else {
		fmt.Printf("myError false\n") // Print
	}

	// As (Assertion Type Casting)
	fmt.Printf("\n--- As ---\n")
	var iErr *innerErr
	if errors.As(outerErr, &iErr) {
		fmt.Printf("innerError ture: %v\n", iErr.Msg) // Print
	} else {
		fmt.Printf("innerError false\n")
	}
	var mErr *myErr
	if errors.As(outerErr, &mErr) {
		fmt.Printf("myError true: %v\n", mErr.Msg)
	} else {
		fmt.Printf("myError false\n") // Print
	}
}
