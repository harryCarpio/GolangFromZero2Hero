/*
Statement
Create a calculator struct with a result attribute and with the most common methods: 
	- Add
	- Subtract
	- Multiply
	- Divide
Plus: Handle possible error, Divide function won’t be able to handle second parameter with value zero, instead function should panic and print an alert

Topics to Practice:
function, return value, panic, defer
*/
package main

import (
	"fmt"
)

type Calculator struct{
	result	int
}

func (c *Calculator) Add(x int, y int) int{
	c.result = x + y
	return c.result
}

func (c *Calculator) Substract(x int, y int){
	c.result = x - y
}

func (c *Calculator) Multiply(x int, y int){
	c.result = x * y
}

func (c *Calculator) Divide(x int, y int){
	defer func() {
        if err := recover(); err != nil {
            fmt.Printf("Alert: %s. Cannot calculate %d/%d\n", err, x, y)
		}else{
			c.result = x / y
		}
    }()

	if y == 0 {	
		panic("Zero division error")
	}
}

func main(){

	calc := Calculator{
	}
	
	fmt.Println("--- Add ---")
	addResult := calc.Add(5, 6)
	fmt.Printf("5 + 6 = %d\n", addResult)

	fmt.Println("--- Subs ---")
	calc.Substract(40, 74)
	fmt.Printf("40 - 74 = %d\n", calc.result)

	fmt.Println("--- Multiply ---")
	calc.Multiply(8, 9)
	fmt.Printf("8 * 9 = %d\n", calc.result)

	fmt.Println("--- Divide ---")
	calc.Divide(10, 0)

	fmt.Println("--- Divide ---")
	calc.Divide(10, 2)
	fmt.Printf("10 / 2 = %d\n", calc.result)


}
