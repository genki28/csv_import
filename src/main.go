package main

import (
	"log"
	"net/http"

	// "strings"

	"csv_import/controllers"
	"csv_import/multiThread"
)

func ManyImport(w http.ResponseWriter, r *http.Request) {
	if (r.Method != http.MethodPost) {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		w.Write([]byte("POSTでこい"))
		return
	}
}

func main() {
	http.HandleFunc("/singleImport", controllers.Single)
	http.HandleFunc("/manyImport", ManyImport)
	http.HandleFunc("/multi", multiThread.Multi)
	http.HandleFunc("/single", multiThread.Hoge)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}