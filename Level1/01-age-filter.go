package main

import (
	"fmt"
)

func myFunc(num1 int, num2 int, slc []int) []int{
	var response = []int{}
	
	for i := range slc {
		if slc[i] > num1 && slc[i] < num2 {
			response = append(response, slc[i])
		}
	}
	
	return response
}

func main() {
	num1:=4
	num2:=11
	var slc = []int{2, 3, 5, 7, 11, 13}
	
	fmt.Println(myFunc(num1, num2, slc))
}
