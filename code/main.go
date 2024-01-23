package main

import (
	"errors"
	"fmt"
)

func main() {
	printValue("Welcome to GO-land!")

	// Variable
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

	learnArray()
	learnSlice()
	learnMap()
	learnString()
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

func learnArray() {
	// Ways of Declaring Arrays -
	var intArr1 [3]int // Filled with Default value of datatype
	intArr2 := [3]int{1, 2, 3}
	intArr3 := [...]int{1, 2, 3, 4, 5} // Without specifying exact length

	fmt.Println(intArr1)
	fmt.Println(intArr2)
	fmt.Println(intArr3)

	// Iterate Array
	for val := range intArr3 {
		fmt.Println(val)
	}

	for i, val := range intArr3 {
		fmt.Printf("Index = %v, Value = %v \n", i, val)
	}
}

func learnSlice() {
	// Slice is just a wrapper around Array
	var intSlice []int = []int{1, 2, 3, 4}
	fmt.Printf("Length = %v, Capacity = %v \n", len(intSlice), cap(intSlice))

	intSlice = append(intSlice, 5)
	fmt.Printf("Length = %v, Capacity = %v \n", len(intSlice), cap(intSlice))

	// Appending multiple elements to slice
	intSlice2 := []int{6, 7}
	intSlice = append(intSlice, intSlice2...)
	fmt.Printf("Length = %v, Capacity = %v \n", len(intSlice), cap(intSlice))

	// Iterate Slice
	for i, val := range intSlice {
		fmt.Printf("Index = %v, Value = %v \n", i, val)
	}
}

func learnMap() {
	// Empty Map - map1 = map[]
	// key type -> string, value type -> int
	var map1 map[string]int = make(map[string]int)
	fmt.Println(map1)

	//Initialized Map
	map2 := map[string]int{"Neeraj": 11, "AI": 10}
	fmt.Println(map2)
	fmt.Println(map2["Neeraj"])

	// Map returns default type value if key doesn't exist
	fmt.Println(map2["Common Sense"])
	map2["Common Sense"] = 101

	// Map always return a bool alongwith the value of a key
	// Book value = true, if key exists
	// 			  = false, if key doesn't exist
	var val, ok = map2["Common Sense"]
	if ok {
		fmt.Printf("Value Exists for the given Key : %v \n", val)
	} else {
		fmt.Printf("Key doesn't exist in the map \n")
	}

	// Delete value from map
	delete(map2, "AI")

	// Iterate Array
	for name, val := range map2 {
		fmt.Printf("Key = %v, Value = %v \n", name, val)
	}
}

func learnString() {
	// String is immutable in Go
	var str string = "neeraj"
	fmt.Printf("value = %v, type = %T \n", str[0], str[0])
}
