/*
===  Interface Switch ===

Statement
Create a function that will receive an interface and it will print the interface value with a specific message based on the type.

Topics to Practice: 
Interfaces, type assertion, switch
*/
package main

import (
	"fmt"
)

func InterfaceFunc(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Previous number of %v is %v\n", v, v-1)
	case bool:
		fmt.Printf("This is a %t boolean\n", v)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	InterfaceFunc(21)
	InterfaceFunc("hello")
	InterfaceFunc(false)
	InterfaceFunc(12.2)
}