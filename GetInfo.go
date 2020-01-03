package FPatephone

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type GetInfoStruct struct {
	Book    Book `json:"book"`
	Success bool `json:"success"`
}

func (api *API) GetInfo(id int) GetInfoStruct {
	url := "https://api.patephone.com/client-api/books/" + strconv.Itoa(id)
	fmt.Println(url)
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
	//fmt.Printf("content: %s\n", string(contents))
	var infoRes GetInfoStruct
	err = json.Unmarshal(contents, &infoRes)
	check(err)
	//fmt.Println("id=", infoRes.Book.ID, " | ", infoRes.Book.Title, ". ", infoRes.Book.Authors[1].FirstName, " ", infoRes.Book.Authors[1].LastName, "\n", infoRes.Book.Description, infoRes.Book.Duration, infoRes.Book.FileSize, infoRes.Book.PreviewStreamURL)
	return infoRes
}
