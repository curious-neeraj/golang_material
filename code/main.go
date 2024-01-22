package main

import (
	"errors"
	"fmt"
)

func main() {
	printValue("Welcome to GO-land!")

	var value1 int = 25
	var value2 int = 10

	// Direct Assigning value - use of :=
	// This is equivalent to :
	// var quotient, remainder int = divide(value1, value2)
	quotient, remainder, err := divide(value1, value2)
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("Quotient = %v\n", quotient)
	} else {
		fmt.Printf("Quotient = %v, Remainder = %v\n", quotient, remainder)
	}
}

func printValue(value string) {
	fmt.Println(value)
}

func divide(numerator int, denominator int) (int, int, error) {

	// Type = error, Defualt Value = nil
	var err error

	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}

	// Division results in whole number rounded down
	quotient := numerator / denominator
	remainder := numerator % denominator
	return quotient, remainder, err
}
