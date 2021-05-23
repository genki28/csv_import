package controllers

import (
	"net/http"
	"encoding/csv"
	"encoding/json"
	"io"
	"flag"
	"log"
	"fmt"
)

type response map[string]interface{}
var jsonSeq *bool = flag.Bool("jsonSeq", false, "output JSON text sequence format")

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
	results := []response{}
	// read header
	header, err := reader.Read()
	if err == io.EOF {
		return
	}

	for {
		// read csv body
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error: csv Read fail: %v\n", err)
		}

		responseData := make(response)
		for i := range line {
			responseData[header[i]] = string(line[i])
		}

		if *jsonSeq {
			// output immediately
			body, err := json.Marshal(responseData)
			if err != nil {
				log.Fatalf("Error: json.Marshal fail: Input: %v, Message: %v", results, err)
			}
			fmt.Printf("%s\n", body)
		} else {
			results = append(results, responseData)
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