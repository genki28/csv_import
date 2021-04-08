package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"fmt"
	"net/http"
	"log"
	"flag"
	// "strings"
)

type JSON map[string]interface{}
var jsonSeq *bool = flag.Bool("jsonSeq", false, "output JSON text sequence format")

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func Single(w http.ResponseWriter, r *http.Request) {
	if (r.Method != http.MethodPost) {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		w.Write([]byte("POSTでこい"))
		return
	}
	file, _, err := r.FormFile("file")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	results := []JSON{}
	// read header
	header, err := reader.Read()
	if err == io.EOF {
		return
	}

	for {
		// read csv body
		rows, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error: csv Read fail: %v\n", err)
		}

		jsonData := make(JSON)
		for i := range rows {
			jsonData[header[i]] = string(rows[i])
		}

		if *jsonSeq {
			// output immediately
			body, err := json.Marshal(jsonData)
			if err != nil {
				log.Fatalf("Error: json.Marshal fail: Input: %v, Message: %v", results, err)
			}
			fmt.Printf("%s\n", body)
		} else {
			results = append(results, jsonData)
		}
	}

	if !*jsonSeq {
		// Save the JSON file
		json, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			log.Fatalf("Error: json.Marshal fail: Input: %v, Message: %v", results, err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
		// fmt.Fprintln(w, json)
	}
}

func main() {
	http.HandleFunc("/singleImport", Single)
	http.HandleFunc("/", Hello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}