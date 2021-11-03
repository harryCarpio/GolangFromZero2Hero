package main

import "fmt"

func sum(s int, c chan int) {
	sum := 0
	sum = sum + s
	c <- sum // send sum to c
}

func main() {

	c := make(chan int)
	go sum(1, c)
	go sum(2, c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
