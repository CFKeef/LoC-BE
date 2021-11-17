package main

import (
	"api/model"
	"api/util"
	"encoding/json"
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(fetch)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/search", util.CORS(handleSearch))
	log.Print("Running")
	log.Fatal(http.ListenAndServe(":80", mux))
}
