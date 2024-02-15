package main

import "fmt"

/*
Write a sumInt function that takes a variable number of int arguments,
and returning the count of arguments and their sum.
*/

func sumInt(n ...int) (c, s int) {
	for _, e := range n {
		c++
		s += e
	}
	return
}

func main() {
	fmt.Println(sumInt(1, 2, 3, 4, 5))
}
