package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type config struct {
	nextUrl *string
	prevUrl *string
}

func showMap(cfg *config) error {
	url := *cfg.nextUrl
	req, _ := http.Get(url)

	var respnse ApiResponse

	decoder := json.NewDecoder(req.Body)
	_ = decoder.Decode(&respnse)

	cfg.nextUrl = respnse.Next
	cfg.prevUrl = respnse.Previous

	for _, nam := range respnse.Results {
		fmt.Println(nam.Name)
	}
	return nil

}

func prevMap(cfg *config) error {
	url := *cfg.prevUrl
	req, _ := http.Get(url)

	var respnse ApiResponse

	decoder := json.NewDecoder(req.Body)
	_ = decoder.Decode(&respnse)

	cfg.nextUrl = respnse.Next
	cfg.prevUrl = respnse.Previous

	for _, nam := range respnse.Results {
		fmt.Println(nam.Name)
	}
	return nil

}
