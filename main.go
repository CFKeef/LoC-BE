package main

import (
	"api/model"
	"api/util"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handleSearch(w http.ResponseWriter, req *http.Request) {
	if len(req.URL.Query()) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)

		res := model.Error{
			Code:    "missing_query",
			Message: "Query wasn't provided",
		}

		json.NewEncoder(w).Encode(res)

		return
	}

	fetch, err := util.HandleFetch(req.URL.Query())

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)

		res := model.Error{
			Code:    "api_error",
			Message: "Something went wrong :(",
		}

		json.NewEncoder(w).Encode(res)

		log.Println(err)

		return
	}

	fmt.Println(fetch)
	w.WriteHeader(200)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/search", util.CORS(handleSearch))
	log.Fatal(http.ListenAndServe(":8000", mux))
}
