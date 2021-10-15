/*
Simple Queue
Statement
Create a Queue implementation. To build it we will need a slice of int, and an enqueue function to add an int into the 
slice and dequeue function to remove the first element of the slice.


Topics to Practice:
functions, slice, append, data structure
*/
package main

import(
	"fmt"
)

type Queue struct {
	slice []int
}

func (q *Queue) Enqueue(item int){
	q.slice = append(q.slice, item)
}

func (q *Queue) Dequeue(){
	q.slice[0] = 0
	q.slice = q.slice[1:]
}

func main(){
	queue := Queue{}
	
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)

	fmt.Println(queue)

	queue.Dequeue()
	fmt.Println(queue)

	queue.Dequeue()
	fmt.Println(queue)
}