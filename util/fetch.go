package util

import (
	"api/config"
	"api/model"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// Handles the fetching from the library api
func HandleFetch(q url.Values) (model.LoCAPIResponse, error) {
	req, err := http.NewRequest("GET", config.BasePath, nil)

	if err != nil {
		return model.LoCAPIResponse{}, err
	}

	// The request needs to have fo=json as a query param of it to get back a json response
	q.Add("fo", "json")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return model.LoCAPIResponse{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return model.LoCAPIResponse{}, err
	}

	var converted model.LoCAPIResponse

	json.Unmarshal(body, &converted)

	return converted, nil
}
