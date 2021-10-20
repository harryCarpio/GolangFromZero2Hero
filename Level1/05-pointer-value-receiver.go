/*
Pointer Receiver and Value Receiver

Statement:
Define a struct name MySlice1, define a slice of indexes as []int. Define addIndex method for the struct which append an int to indexes ( pass int as parameter), make sure it is defined as Value Receiver.
Define another struct MySlice2, make everything the same as MySlice2 unless define it as Pointer Receiver.
Create a main method, and define each of structs and use addIndex few times for each. Print value of indexes for each and compare the result.

Topics to Practice:
Pointer receiver and Value receiver
*/
package main

import(
	"fmt"
)

type MySlice1 struct{
	indexes []int
}

func (s1 *MySlice1) addIndex(index int){
	s1.indexes = append(s1.indexes, index)
}

type MySlice2 struct{
	indexes []int
}

func (s2 *MySlice2) addIndex(index *int){
	s2.indexes = append(s2.indexes, *index)
}

func main(){
	slc1 := MySlice1{}
	slc1.indexes = []int{}

	slc2 := MySlice2{}
	slc2.indexes = []int{}

	var idx = 1

	slc1.addIndex(idx)
	slc2.addIndex(&idx)

	fmt.Println(slc1)
	fmt.Println(slc2)
}