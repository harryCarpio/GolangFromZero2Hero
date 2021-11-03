package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Hero struct {
	Name  string
	Color string
}

func GetHero(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	heroName := params["hname"]

	var resultHero Hero

	for _, h := range heroes {
		if h.Name == heroName {
			resultHero = h
		}
	}

	w.Header().Set("Content-Type", "application/json")

	if resultHero == (Hero{}) {
		json.NewEncoder(w).Encode(nil)
	} else {
		json.NewEncoder(w).Encode(resultHero)
	}
}

var heroes []Hero

func main() {
	router := mux.NewRouter()

	heroes = append(heroes, Hero{Name: "Ironman", Color: "Red"})
	heroes = append(heroes, Hero{Name: "Batman", Color: "Black"})
	heroes = append(heroes, Hero{Name: "Superman", Color: "Blue"})
	heroes = append(heroes, Hero{Name: "Green Arrow", Color: "Green"})

	router.HandleFunc("/heroes/{hname}", GetHero).Methods("GET")

	server := &http.Server{
		Handler: router,
		Addr:    "localhost:8080",
	}

	log.Fatal(server.ListenAndServe())
}
