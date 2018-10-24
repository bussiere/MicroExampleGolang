package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Todo struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Created   time.Time `json:"createdon"`
}

func main() {

	url := "http://localhost:8080/api/todos"
	var tr Todo
	tr.Name = "titi"
	jsonStr, err := json.Marshal(tr)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record Todo

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	log.Println(record)
	log.Println(record.ID)
	log.Println(record.Name)
}
