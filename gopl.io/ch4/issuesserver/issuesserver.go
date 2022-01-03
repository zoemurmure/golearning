package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/report", generateReport)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func generateReport(w http.ResponseWriter, r *http.Request) {

}
