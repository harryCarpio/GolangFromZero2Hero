/*
=== Simple Inventory ===
Statement
Create a Product struct with an attribute ID and Name, and an Inventory struct that will contain a map[string]Products. 
Create a method with a pointer of Inventory as a receiver and It will expect a product to be added into the map of products. 
Required: Add validations into the Add method, to check that the ID can’t be empty and also not duplicated ID

Topics to Practice:
struct, composition, methods, pointer, error handling
*/
package main

import(
	"fmt"
	"strings"
)

type Product struct{
	id string
	name string
}

type Inventory struct{
	products map[string]Product
}

func (i *Inventory) AddProduct(pr Product){
	defer func(){
		if err := recover(); err != nil {
            fmt.Printf("Alert: %s\n", err)
		}else{
			i.products[pr.id] = pr
		}
	}()

	_, keyExists := i.products[pr.id]
	if keyExists {
		panic("Product ID "+pr.id+" is already registered. Product "+pr.name+" not registered.")
	}

	valSpaces := strings.Trim(pr.id, " ")
	if len(valSpaces) == 0 {
		panic("Product ID is empty.")
	}
}

func main(){
	inventory := Inventory{products: make(map[string]Product)}

	pr1 := Product{id: "P001", name: "Oreo"}
	inventory.AddProduct(pr1)
	fmt.Println(inventory)

	pr2 := Product{id: "P002", name: "KitKat"}
	inventory.AddProduct(pr2)
	fmt.Println(inventory)

	pr3 := Product{id: "P002", name: "Froyo"}
	inventory.AddProduct(pr3)
	fmt.Println(inventory)

	pr4 := Product{id: "", name: "Froyo"}
	inventory.AddProduct(pr4)
	fmt.Println(inventory)

	pr5 := Product{id: "   ", name: "Froyo"}
	inventory.AddProduct(pr5)
	fmt.Println(inventory)

	pr6 := Product{id: "P003", name: "Froyo"}
	inventory.AddProduct(pr6)
	fmt.Println(inventory)
}