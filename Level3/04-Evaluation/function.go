package main

import (
	"fmt"
)

func filterByAgeRange(fromAge, toAge int, ages []int) []int{
	response := make([]int, 0)
	for _,ageToCheck := range ages {
		if fromAge <= ageToCheck && ageToCheck <= toAge  {
			response = append(response, ageToCheck)
		}
	}
	return response 
}

func main(){
	fmt.Println("Function")
}

/*
You can use the following function to Deep compare arrays
func Equal(a, b []int) bool {
	if len(a) != len(b) {
			return false
	}
	for i, v := range a {
		if v != b[i] {
				return false
		}
	}
	return true
}
*/