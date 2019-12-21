package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestCart_SaveObject(t *testing.T) {
	type args struct {
		item Item
	}
	tests := []struct {
		name string
		cart *Cart
		args args
	}{
		{
			name: "Test that save to the struct the item",
			cart: &Cart{},
			args: args{
				item: Item{
					Name:  "Blue dress",
					Price: 321.23,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cart.saveObject(tt.args.item)
			if tt.args.item != tt.cart.Item[0] {
				t.Errorf("got %v want %v", tt.args.item, tt.cart.Item)
			} 
			log.Println(tt.cart.Item)
		})
	}
}

func Test_requestPOSTAPI(t *testing.T) {
	url := "http://localhost:8181/addItem"
	payload := strings.NewReader("{\"name\":\"test1\",\"price\":123.00}")
	req, _ := http.NewRequest("POST", url, payload)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func Test_requestGETAPI(t *testing.T) {
	url := "http://localhost:8181/getItems"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
