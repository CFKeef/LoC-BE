package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleSearch(w http.ResponseWriter, req *http.Request) {
	fmt.Print(req.URL.Query())
	w.WriteHeader(200)

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/search", handleSearch)

	log.Fatal(http.ListenAndServe(":8000", mux))
}
