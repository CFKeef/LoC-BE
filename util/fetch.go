package util

import (
	"api/config"
	"api/model"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

func HandleFetch(q url.Values) (model.LoCAPIResponse, error) {
	req, err := http.NewRequest("GET", config.BasePath, nil)

	if err != nil {
		return model.LoCAPIResponse{}, err
	}

	q.Add("fo", "json")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.LoCAPIResponse{}, err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return model.LoCAPIResponse{}, err
	}

	var converted model.LoCAPIResponse

	json.Unmarshal(body, &converted)

	return converted, nil
}
