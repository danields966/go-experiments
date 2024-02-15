package main

import "fmt"

/*
Two channels are given. The first channel receives numbers.
It's necessary that the numbers is receiving in one goroutine.
In the second goroutine they were squared and the result sent to the second channel.
*/

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := range ch1 {
			ch2 <- i * i
		}
	}()

	for j := range ch2 {
		fmt.Println(j)
	}
}
