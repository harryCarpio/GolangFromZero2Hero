/*
=== Shape with Interfaces ===

Statement
Create a Rectangle struct and a Circle struct, both struct should have two Method:Area() Perimeter()
And then create a Shape interface for those struct. After that create a function that will receive a Shape interface as parameter and will execute the Area() and the Perimeter() from each struct. 

Topics to Practice:
Interfaces, struct, method, function, math pkg
*/
package main

import(
	"fmt"
	"math"
)

type Rectangle struct{
	width float64
	height float64
}

func (r Rectangle) Area() float64{
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return (2*r.width) + (2*r.height)
}

type Circle struct{
	radius float64
}

func (c Circle) Area() float64{
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) Perimeter() float64{
	return 2 * math.Pi * c.radius
}

type Shape interface{
	Area() float64
	Perimeter() float64
}

func measure(s Shape){
	switch shape := s.(type){
	case Circle:
		fmt.Printf("\n- Circle with radius %v measures \n", shape.radius)
	case Rectangle:
		fmt.Printf("\n- Rectangle with width %v and heigth %v measures\n", shape.width, shape.height)
	default:
		fmt.Println("I don't know this shape")
	}
	fmt.Printf("Area %v\n",s.Area())
	fmt.Printf("Perimeter %v\n",s.Perimeter())
}

func main(){
	r1 := Rectangle{2, 3.5}
	c1 := Circle{4}

	measure(r1)
	measure(c1)
}

