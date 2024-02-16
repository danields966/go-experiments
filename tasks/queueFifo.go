package main

/*
Implement FIFO queue using linked lists.
There should be three functions:
Push (adding an element)
Pop (removing an element and returning it)
printQueue (print the queue on one line without spaces)
*/

import (
	"container/list"
	"fmt"
)

func Push(elem interface{}, queue *list.List) {
	queue.PushBack(elem)
}

func Pop(queue *list.List) interface{} {
	e := queue.Front()
	queue.Remove(e)
	return e.Value
}

func printQueue(queue *list.List) {
	for temp := queue.Front(); temp != nil; temp = temp.Next() {
		fmt.Print(temp.Value)
	}
}

func main() {
	queue := list.New()
	Push(1, queue)
	Push(2, queue)
	Push(3, queue)
	printQueue(queue) // 123
	fmt.Println()
	Pop(queue)
	printQueue(queue) // 23
}
