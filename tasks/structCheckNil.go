package main

import "fmt"

type P struct {
	a, b int
}

func main() {
	var p *P = nil
	if p == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}

	var p2 interface{} = (*P)(nil)
	if p2 == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}
