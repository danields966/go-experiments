package main

/*
Write a goroutine with the following signature:
func calculator(arguments <-chan int, done <-chan struct{}) <-chan int
This function takes two read-only channels and returns an output read-only channel.
The arguments channel will receive a series of numbers, and the done channel will receive a signal to return result.
When the completion signal is received, the function must send the sum the numbers to the output channel.
The calculator function must be non-blocking and return the control immediately.
*/

import "fmt"

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	outputChannel := make(chan int)
	go func() {
		s := 0
		for {
			select {
			case v := <-arguments:
				s += v
			case <-done:
				outputChannel <- s
				close(outputChannel)
				return
			}
		}
	}()
	return outputChannel
}

func main() {
	arguments := make(chan int)
	done := make(chan struct{})
	result := calculator(arguments, done)
	for i := 0; i < 10; i++ {
		arguments <- i
	}
	close(done)
	fmt.Println(<-result)
}
