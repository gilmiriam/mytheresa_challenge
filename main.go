package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type server struct{}

//Item defines the item input
type Item struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

//Cart defines the entire cart
type Cart struct {
	Item []Item
}

var cart *Cart = &Cart{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(cart)
	case "POST":
		item := Item{}
		err := json.NewDecoder(r.Body).Decode(&item)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		cart.saveObject(item)
		json.NewEncoder(w).Encode(item)
		w.Write([]byte(`{"Status": "Item Added"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"Response": "Not Found"}`))
	}
}

func (saveCart *Cart) saveObject(item Item) {
	saveCart.Item = append(saveCart.Item, item)
}

func main() {
	s := &server{}
	http.Handle("/addItem", s)
	http.Handle("/getItems", s)
	log.Fatal(http.ListenAndServe(":8181", nil))
}
