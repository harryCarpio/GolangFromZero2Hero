/*
Statement 3
Create a channel of int and then create a goroutine to add a value to the channel and then 
print the channel value in the main function

Topics to Practice: 
goroutine, channel, function
*/
package main

import(
	"fmt"
)

func main(){
	num := 0
	ints := make(chan int)

	go func(){
		num = num + 1
		ints <- num
	}()

	chanVal := <- ints
	fmt.Println(chanVal)
}