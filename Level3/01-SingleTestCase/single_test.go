/*
>>> Single Test case

Statement
Create a Testing function to check the behavior of the following function. 
If the function returns a different value from the expected one, return an error specifying the test case. 

Topics to Practice:
testing, function, conditions, error
*/
package main

import(
	"fmt"
	"testing"
)

func IntMin(a, b int) int{
	if a < b{
		return a
	}
	return b
}

func TestIntMin(t *testing.T){
	var minValue int
	minValue = IntMin(3, 4)
	if minValue != 3 {
		t.Error("Expected 3, got ", minValue)
	}
}

func main(){
	fmt.Println("Single test")
}