package main

/*
#cgo LDFLAGS: -L./calculator -lcalculator
#include "calculator/calculator.h"
*/
import "C"
import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Read input values from command line arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: main <value1> <value2>")
		os.Exit(1)
	}

	arg1, err1 := strconv.Atoi(os.Args[1])
	arg2, err2 := strconv.Atoi(os.Args[2])

	if err1 != nil || err2 != nil {
		fmt.Println("Invalid input. Please provide valid integers.")
		os.Exit(1)
	}

	fmt.Printf("Invoking C++ library...\n")
	// Create an instance of the Calculator class
	calculator := C.Calculator{}

	// Call the add method
	sum := C.int(calculator.add(C.int(arg1), C.int(arg2)))
	fmt.Printf("Sum: %d\n", sum)

	// Call the sub method
	diff := C.int(calculator.sub(C.int(arg1), C.int(arg2)))
	fmt.Printf("Difference: %d\n", diff)
}
