package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type server struct{}

type Item struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Cart struct {
	Item []Item
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	// case "GET":
	// 	w.WriteHeader(http.StatusOK)
	// 	keys, ok := r.URL.Query()["key"]

	// 	if !ok || len(keys[0]) < 1 {
	// 		log.Println("Url Param 'key' is missing")
	// 		return
	// 	}

	// 	key := keys[0]

	// 	w.Write([]byte(`{"Input Slice": "TEST"}`))
	case "POST":
		var item Item
		err := json.NewDecoder(r.Body).Decode(&item)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		fmt.Fprintf(w, item.Name)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"Response": "Not Found"}`))
	}
}

func main() {
	s := &server{}
	http.Handle("/addItem", s)
	log.Fatal(http.ListenAndServe(":8181", nil))
}
