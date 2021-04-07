package main

import (
	"encoding/csv"
	"io"
	"fmt"
	"net/http"
	"log"
	"strings"
)

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

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		output := strings.Join(line[:], " for ") + "\n"
		fmt.Fprintf(w, output)
	}
}

func main() {
	http.HandleFunc("/singleImport", Single)
	http.HandleFunc("/", Hello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}