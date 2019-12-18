package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type server struct{}

type cart struct {
	name  string
	price float64
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
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}

		log.Println(string(body))
		// results = append(results, string(body))

		fmt.Fprint(w, "POST done")
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
