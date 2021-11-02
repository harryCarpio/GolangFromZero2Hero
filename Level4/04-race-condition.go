/*
Statement 4
Create a function that will increase a number and that function will be
executed by a goroutine inside a for loop (x1000 times). To avoid race conditioning,
implement the sync.Mutex and Lock and Unlock inside the increase() function.
Note: Add a time.Sleep() to be able to see the final n

Topics to Practice:
goroutine, sync.mutex, function, for loop, defer, time pkg
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeIncrementer struct {
	mu     sync.Mutex
	number int
}

func (i *SafeIncrementer) IncreaseNumber() {
	i.mu.Lock()
	i.number++
	i.mu.Unlock()
}

func main() {
	si := SafeIncrementer{number: 0}
	n := 0

	for n < 1000 {
		go si.IncreaseNumber()
		n++
	}

	time.Sleep(time.Second * 3)
	fmt.Println(si.number)
}
