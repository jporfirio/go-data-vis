package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
)

type DataFile struct {
	Data [][2]any `json:"data"`
}

func main() {
	file, err := os.ReadFile("data/data.json")
	if err != nil {
		log.Fatalf("Error reading file: %v\n", err)
	}

	var data DataFile
	if err := json.Unmarshal(file, &data); err != nil {
		log.Fatalf("Error decoding JSON: %v\n", err)
	}

	component := Chart(data.Data)

	http.Handle("/", templ.Handler(component))

	fmt.Printf("Listening on port :%d\n", 8080)

	http.ListenAndServe(":8080", nil)
}
