package main

/*
Write a goroutine with the following signature:
func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int)
It must do the following calculations for the n times:
* read one number from each of two channels in1 and in2, let's call them x1 and x2.
* calculate fn(x1) + fn(x2)
* write the result to out channel
The function must be non-blocking and return control immediately.
The fn function can run for a long time, and it must be called asynchronously.
*/

import (
	"fmt"
	"sync"
	"time"
)

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		defer close(out)

		wg := new(sync.WaitGroup)
		mu := new(sync.Mutex)

		l1 := make([]int, 0, n)
		l2 := make([]int, 0, n)
		m := make(map[int]int)

		for i := 0; i < n*2; i++ {
			select {
			case val := <-in1:
				l1 = append(l1, val)
			case val := <-in2:
				l2 = append(l2, val)
			}
		}

		if len(l1) != n || len(l2) != n {
			fmt.Println("Input channels did not receive expected number of elements.")
			return
		}

		for i := 0; i < n; i++ {
			val1 := l1[i]
			val2 := l2[i]

			wg.Add(2)
			go func() {
				defer wg.Done()
				r := fn(val1)
				mu.Lock()
				m[val1] = r
				mu.Unlock()
			}()
			go func() {
				defer wg.Done()
				r := fn(val2)
				mu.Lock()
				m[val2] = r
				mu.Unlock()
			}()
		}

		wg.Wait()
		for i := 0; i < n; i++ {
			out <- m[l1[i]] + m[l2[i]]
		}
	}()
}

func main() {
	const n = 20

	fn := func(x int) int {
		time.Sleep(time.Second)
		return x * 2
	}
	in1 := make(chan int)
	in2 := make(chan int)
	out := make(chan int, n)

	go merge2Channels(fn, in1, in2, out, n)

	for i := 0; i < n; i++ {
		in1 <- i
		in2 <- i
	}

	close(in1)
	close(in2)

	orderFail := false
	evenFail := false
	prev := -1

	for i := 0; i < n; i++ {
		c := <-out
		if c%2 != 0 {
			evenFail = true
		}
		if prev >= c && i != 0 {
			orderFail = true
		}
		prev = c
		fmt.Println(c)
	}

	if orderFail {
		fmt.Println("Order violation detected.")
	}
	if evenFail {
		fmt.Println("Odd numbers detected.")
	}
}
