/*
A store sells two types of products: books and games. 
Each product has the fields name (string) and price (float). 
Define the following functionalities:
	Product has the following methods:
		A method to print the information of each product (type, name, and price)
		A method to apply a discount ratio to the price
The store should be able to apply custom discounts based on the type of product: 
	10% discount for books and 20% discount for games.
*/
package main

import(
	"fmt"
)

type Product interface{
	PrintInfo()
	ApplyDiscount()
}

type Books struct{
	name string
	price float64
}

func (b *Books) PrintInfo(){
	fmt.Printf("Type: Book, Name: %q, Price: %v\n", b.name, b.price)
}

func (b *Books) ApplyDiscount(){
	b.price = b.price * 0.9
}

type Games struct{
	name string
	price float64
}

func (g *Games) PrintInfo(){
	fmt.Printf("Type: Game, Name: %q, Price: %v\n", g.name, g.price)
}

func (g *Games) ApplyDiscount(){
	g.price = g.price * 0.8
}

func main(){
	b1 := Books{name:"Harry Potter", price:10.5}
	g1 := Games{name:"Age of empires", price:40.99}

	b1.PrintInfo()
	g1.PrintInfo()

	b1.ApplyDiscount()
	g1.ApplyDiscount()
	fmt.Println("After discounts")
	b1.PrintInfo()
	g1.PrintInfo()
}