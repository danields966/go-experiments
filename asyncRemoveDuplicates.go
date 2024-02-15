package main

/*
Write a goroutine with the following signature:
func removeDuplicates(inputStream, outputStream chan string)
This goroutine must remember the previous value from inputStream and send the value to the outputStream
only if it's different from previous value.
The removeDuplicates function must be non-blocking and return the control immediately.
*/

import (
	"fmt"
	"time"
)

func removeDuplicates(inputStream, outputStream chan string) {
	defer close(outputStream)
	var p string
	for i := range inputStream {
		if i != p {
			outputStream <- i
			p = i
		}
	}
}

func printChannel(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
	}
}

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)
	defer close(inputStream)

	go removeDuplicates(inputStream, outputStream)
	go printChannel(outputStream)

	for _, i := range "11233445666" {
		inputStream <- string(i)
	}
	time.Sleep(time.Second * 1)
}
