package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type server struct{}

var cart *Cart = &Cart{}

//Item defines the item input
type Item struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

//Cart defines the entire cart
type Cart struct {
	Item []Item
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		js, err := json.Marshal(cart)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(js)
	case "POST":
		item := Item{}
		err := json.NewDecoder(r.Body).Decode(&item)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		cart.SaveObject(item)

	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"Response": "Not Found"}`))
	}
}

//SaveObject save the object into Cart struct
func (saveCart *Cart) SaveObject(item Item) {
	saveCart.Item = append(saveCart.Item, item)
	log.Println(saveCart)
}

func main() {
	s := &server{}
	http.Handle("/addItem", s)
	http.Handle("/getItems", s)
	log.Fatal(http.ListenAndServe(":8181", nil))
}
