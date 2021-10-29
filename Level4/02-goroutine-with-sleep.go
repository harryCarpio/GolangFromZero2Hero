/*
Statement 2
Create a function goroutine that will execute an anonymous function to just print the number “1”, 
in the main function print the number “0” and also add a time.Sleep() to wait 2 seconds.

Topics to Practice: 
goroutine, function, time pkg
*/
package main

import(
	"fmt"
	"time"
)

func main(){
	go func(msg string) {
        fmt.Println(msg)
    }("1")

	fmt.Println("0")
	time.Sleep(time.Second * 2)
}