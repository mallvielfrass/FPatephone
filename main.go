package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
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

type Paging struct {
	Page     int `json:"page"`
	Limit    int `json:"limit"`
	Offset   int `json:"offset"`
	OffsetID int `json:"offset_id"`
	Count    int `json:"count"`
	CountAll int `json:"count_all"`
}
type Request struct {
	Success bool   `json:"success"`
	Book    []Book `json:"books"`
	Paging  Paging `json:"paging"`
}

func Select() int {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("---------------------")
	fmt.Println("Select book, please!")
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	z, _ := strconv.Atoi(string(text))
	//	fmt.Println(z)

	return z
}
func Init() {

}
func main() {
	client := &http.Client{}
	req, err := http.NewRequest(
		"GET", "https://api.patephone.com/client-api/search/book?search_string=time", nil,
	)

	req.Header.Add("X-CLIENT-IDENTIFIER", "patephone_unlim_android")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err %s", err)
		os.Exit(1)
	}
	//fmt.Printf("content: %s\n", string(contents))
	//fil:=string(contents)
	var result Request
	err = json.Unmarshal(contents, &result)
	if err != nil {
		log.Fatalln("f", err)
	}
	//log.Printf("%#v", result)
	n := result.Paging.Count
	for i := 0; i < n; i++ {
		fmt.Println(i, "\b.) ", result.Book[i].ID, " : ", result.Book[i].Title) //9411
	}
	id := Select()
	fmt.Println("you selected:", id, "\b.)", result.Book[id].Title)
	fmt.Println("---------------------")
	GetInfo(result.Book[id].ID)
}
