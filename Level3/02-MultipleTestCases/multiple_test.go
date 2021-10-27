/*
>>> Multiple Test cases

Statement
Create a Testing function to check the behavior of the following function. 
Create a struct with the function parameter and the ‘want’ (expected) value, 
then iterate all the test cases and validate the behavior of the function against each case.

Topics to Practice: 
testing, multiple cases
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

func TestMultipleIntMin(t *testing.T) {
    var tests = []struct {
        a, b int
        want int
    }{
        {0, 1, 0},
        {1, 0, 0},
        {2, -2, -2},
        {0, -1, -1},
        {-1, 0, -1},
    }

	for _, tt := range tests {
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {
            ans := IntMin(tt.a, tt.b)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}

func main(){
	fmt.Println("Mutiple test")
}
