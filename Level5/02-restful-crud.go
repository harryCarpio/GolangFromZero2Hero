/*
Statement 2
- Create a RESTful API implementing all the CRUD Actions, and integrate using gorilla/mux
- Create a Product struct with the following attributes:
	[ID string, Code string, Name string, Price float64]
- Create a ProductInventory struct with the following attributes:
	[Product Product, Quantity int]
Then to simulate a table in memory, create a var inventory as a []ProductInventory
Now, Create the functions [Add, Update, Delete, Get] that will be executed against inventory.
For example: Add will add a new ProductInventory into the inventory ([]ProductInventory)

Topics to Practice:
net/http, RESTful, []byte, log pkg, gorilla/mux
*/
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Product struct {
	ID    string  `json:"ID"`
	Code  string  `json:"Code"`
	Name  string  `json:"Name"`
	Price float64 `json:"Price"`
}

type ProductInventory struct {
	Product  Product
	Quantity int
}

func Add(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var product Product
	json.Unmarshal(reqBody, &product)

	newPi := ProductInventory{Product: product, Quantity: 1}

	inventory = append(inventory, newPi)

	json.NewEncoder(w).Encode(inventory)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

var inventory []ProductInventory

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/inventory", Add).Methods("POST")

	srv := &http.Server{
		Handler: router,
		Addr:    "localhost:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
