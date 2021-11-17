package model

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type LoCAPIResponse struct {
	Pagination struct {
		Current  int    `json:"current"`
		First    string `json:"first,omitempty"`
		From     int    `json:"from"`
		Last     string `json:"last,omitempty"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results string `json:"results"`
		PageList []struct {
			Number int    `json:"number"`
			URL    string `json:"url,omitempty"`
		}
		Total int `json:"total"`
	} `json:"pagination"`
	Results []struct {
		ID       string   `json:"id"`
		ImageURL []string `json:"image_url"`
		Date     string   `json:"date"`
		Title    string   `json:"title"`
		Item     struct {
			CallNumber   []string `json:"call_number,omitempty"`
			Contributors []string `json:"contributors,omitempty"`
		} `json:"item"`
	} `json:"results"`
}
