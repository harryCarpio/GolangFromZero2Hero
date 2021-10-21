/*
=== Musical Player Greetings ===

Statement
Create a Trumpeter struct and a Violinist struct both struts should have a Name attribute and a Greetings() function to present themself. 
Then create a MusicalPlayer interface. In the main function, create a slice with two or more musical players and for each call the Greetings() function. 

Topics to Practice: 
Type Interfaces, struct, method, function, slice, for each loop
*/

package main

import(
	"fmt"
)

type Trumpeter struct{
	Name string
}

func (t Trumpeter) Greetings(){
	fmt.Printf("My name is %q and I'am a Trumpeter\n", t.Name)
}

type Violinist struct{
	Name string
}

func (v Violinist) Greetings(){
	fmt.Printf("My name is %q and I'am a Violinist\n", v.Name)
}

type MusicalPlayer interface{
	Greetings()
}

func main(){
	var t1 = Trumpeter{Name:"Louis Armstrong"}
	var v1 = Violinist{"Vivaldi"}
	var t2 = Trumpeter{"Miles David"}

	musicians := []MusicalPlayer{t1, v1, t2}

	for _, m := range musicians {
		m.Greetings()
	}
}