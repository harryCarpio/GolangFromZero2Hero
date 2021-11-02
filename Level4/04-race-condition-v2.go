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
)

type SafeIncrementer struct {
	mu     sync.Mutex
	number int
}

func (i *SafeIncrementer) IncreaseNumber(wg *sync.WaitGroup) {
	i.mu.Lock()
	i.number++
	i.mu.Unlock()
	wg.Done()
}

func main() {
	si := SafeIncrementer{number: 0}
	n := 0
	iterations := 1000

	var wg sync.WaitGroup
	wg.Add(iterations)

	for n < iterations {
		go si.IncreaseNumber(&wg)
		n++
	}
	wg.Wait()

	fmt.Println(si.number)
}
