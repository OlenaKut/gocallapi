package main

// ev på onsdag = post och put

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Posts []struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Views int    `json:"views"`
}

// https://mholt.github.io/json-to-go/

type Producter []struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Image2      string  `json:"image2"`
	Rating      struct {
		Rate  float64 `json:"rate"`
		Count int     `json:"count"`
	} `json:"rating"`
}

func main() {
	// this is a comment
	//hej hej
	//url := "https://fakestoreapi.com/products"
	url := "http://localhost:3000/posts"

	res, getErr := http.Get(url)
	if getErr != nil {
		log.Fatal(getErr)
	}
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	//fmt.Println(string(body))

	data_obj := Posts{}

	jsonErr := json.Unmarshal(body, &data_obj)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	for _, value := range data_obj {
		//fmt.Println(key)
		fmt.Printf("%v %v %v\n", value.ID, value.Title, value.Views)
	}
}
