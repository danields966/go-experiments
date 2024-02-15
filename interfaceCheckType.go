package main

/*
If we don't know the type of the object inside the interface, then we can use the type switch statement:
*/

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%d multiplied by 2: %d\n", v, v*2)
	case string:
		fmt.Println(v + " Golang!")
	default:
		fmt.Printf("Don't know this type: %T\n", v)
	}
}

func main() {
	do(21)
	do("Hello")
	do(true)
}
