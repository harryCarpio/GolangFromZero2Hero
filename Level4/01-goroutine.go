/*
Statement 1
Create a goroutine that will execute an anonymous function to print “Hello World” and in the main routine print 
“main function”

Topics to Practice: 
goroutine, function, common Issue
*/
package main

import(
	"fmt"
)

func main(){
	go func(msg string) {
        fmt.Println(msg)
    }("Hello World")

	fmt.Println("main function")
}