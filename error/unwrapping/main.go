package main

import (
	"errors"
	"fmt"
)

var (
	innerError = &innerErr{Msg: "innerMsg"}
	myError    = &myErr{Msg: "myMsg"}
)

// Custom errors
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
		return fmt.Errorf("middleError: %w", err)
	}
	return nil
}

func outerFunc() error {
	if err := middleFunc(); err != nil {
		return fmt.Errorf("outerError: %w", err)
	}
	return nil
}

func main() {
	// Get a wrapped error
	outerErr := outerFunc()

	// Unwrap
	fmt.Printf("--- Unwrap ---\n")
	fmt.Printf("unwrap x 0: %v\n", outerErr)
	fmt.Printf("unwrap x 1: %v\n", errors.Unwrap(outerErr))
	fmt.Printf("unwrap x 2: %v\n", errors.Unwrap(errors.Unwrap(outerErr)))

	// Stack
	fmt.Printf("--- Stack ---\n")
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

	// As (Assertion, Type Casting)
	fmt.Printf("\n--- As ---\n")
	var iErr *innerErr
	if errors.As(outerErr, &iErr) {
		fmt.Printf("innerError true: %v\n", iErr.Msg) // Print
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
