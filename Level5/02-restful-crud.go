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
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var product Product
	json.Unmarshal(reqBody, &product)

	newPi := ProductInventory{Product: product, Quantity: 1}

	newProductExists := false

	for idx, pi := range inventory {
		if pi.Product == newPi.Product {
			pi.Quantity++
			inventory[idx] = pi
			newProductExists = true
		}
	}

	if !newProductExists {
		inventory = append(inventory, newPi)
	}

	json.NewEncoder(w).Encode(inventory)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	deleteProductId := params["pid"]

	sliceItemToRemove := -1

	for idx, pi := range inventory {
		if pi.Product.ID == deleteProductId {
			sliceItemToRemove = idx
		}
	}

	if sliceItemToRemove != -1 {
		inventory = removeSliceItem(inventory, sliceItemToRemove)
	}

	json.NewEncoder(w).Encode(inventory)
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var product Product
	json.Unmarshal(reqBody, &product)

	newPi := ProductInventory{Product: product, Quantity: 1}

	params := mux.Vars(r)
	deleteProductId := params["pid"]
	sliceItemToRemove := -1

	for idx, pi := range inventory {
		if pi.Product.ID == deleteProductId {
			sliceItemToRemove = idx
		}
	}

	if sliceItemToRemove != -1 {
		inventory = removeSliceItem(inventory, sliceItemToRemove)
	}

	inventory = append(inventory, newPi)

	json.NewEncoder(w).Encode(inventory)
}

func removeSliceItem(s []ProductInventory, i int) []ProductInventory {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func Get(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(inventory)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
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
	router.HandleFunc("/inventory", Get).Methods("GET")
	router.HandleFunc("/inventory/{pid}", Delete).Methods("DELETE")
	router.HandleFunc("/inventory/{pid}", Update).Methods("PUT")

	srv := &http.Server{
		Handler: router,
		Addr:    "localhost:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
