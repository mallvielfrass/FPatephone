package FPatephone

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Role struct {
	Name string `json:"name"`
	Abbr string `json:"abbr"`
}
type Authors struct {
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Role       Role   `json:"role"`
}
type Images struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}
type Book struct {
	ID               int       `json:"id"`
	Title            string    `json:"title"`
	ShortDescription string    `json:"short_description"`
	Description      string    `json:"description"`
	Publish          string    `json:"publish"`
	PreviewStreamURL string    `json:"preview_stream_url"`
	Rating           int       `json:"rating"`
	Duration         int       `json:"duration"`
	FileSize         int       `json:"file_size"`
	Authors          []Authors `json:"authors"`
	Images           []Images  `json:"images"`
	PriceTier        int       `json:"price_tier"`
	M4bFileSize      int       `json:"m4b_file_size"`
	LastModification int       `json:"last_modification"`
	Mp3PreviewURL    string    `json:"mp3_preview_url"`
	DateUpdated      int       `json:"date_updated"`
	DateCreated      int       `json:"date_created"`
}
type Request struct {
	Success bool   `json:"success"`
	Book    []Book `json:"books"`
	Paging  Paging `json:"paging"`
}
type Paging struct {
	Page     int `json:"page"`
	Limit    int `json:"limit"`
	Offset   int `json:"offset"`
	OffsetID int `json:"offset_id"`
	Count    int `json:"count"`
	CountAll int `json:"count_all"`
}

func (api *API) SearchBook(target string) Request {
	url := "https://api.patephone.com/client-api/search/book?search_string=" + target
	client := &http.Client{}
	req, err := http.NewRequest(
		"GET", url, nil,
	)
	req.Header.Add("X-CLIENT-IDENTIFIER", api.XClient)
	req.Header.Add("User-Agent", api.UserAgent)
	req.Header.Add("X-AUTH-TOKEN", api.XAuthToken)
	req.Header.Add("X-AD-TOKEN", api.XAdToken)
	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	oscheck(err)
	var result Request
	err = json.Unmarshal(contents, &result)
	check(err)
	return result
}
