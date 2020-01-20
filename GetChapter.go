package FPatephone

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Chapters struct {
	Timecode string `json:"timecode"`
	Name     string `json:"name"`
}
type RequestChapter struct {
	Success  bool       `json:"success"`
	Chapters []Chapters `json:"chapters"`
}

func (api *API) GetChapter(target int) RequestChapter {

	url := "https://api.patephone.com/client-api/books/" + strconv.Itoa(target) + "/chapters"
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
	var resultChapter RequestChapter
	fmt.Println("Debug GetChapter")
	fmt.Println(string(contents))
	err = json.Unmarshal(contents, &resultChapter)
	if err != nil {
		log.Println(err)

	}
	return resultChapter
}
