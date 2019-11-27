package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type GetInfoStruct struct {
	Book    Book `json:"book"`
	Success bool `json:"success"`
}

func Parse(date string) {
	var masser = map[int]string{}
	s := strings.Split(date, "segment")
	z := len(s)
	i := 1
	for i < z {
		masser[i-1] = strings.Split(s[i], ".aac")[0]
		i++
	}
	i = i - 1
	fmt.Println(masser)
	fmt.Println(i)

}
func ChlkDir(id int) {
	path := strconv.Itoa(id)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0777)
	}
}
func Stream(url string) {
	//	/stream/hls/preview/13613/fileList.m3u8
	fmt.Println(url)
	urlx := "https://n01.cd.ru.patephone.com" + url
	fmt.Println(urlx)
	client := &http.Client{}
	req, err := http.NewRequest(
		"GET", urlx, nil,
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
//	fmt.Println(string(contents))
	Parse(string(contents))
	

}
func GetInfo(id int) {
	url := "https://api.patephone.com/client-api/books/" + strconv.Itoa(id)
	fmt.Println(url)
	client := &http.Client{}
	req, err := http.NewRequest(
		"GET", url, nil,
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
	fmt.Printf("content: %s\n", string(contents))
	var infoRes GetInfoStruct
	err = json.Unmarshal(contents, &infoRes)
	if err != nil {
		log.Fatalln("f", err)
	}

	fmt.Println("id=", infoRes.Book.ID, " | ", infoRes.Book.Title, ". ", infoRes.Book.Authors[1].FirstName, " ", infoRes.Book.Authors[1].LastName, "\n", infoRes.Book.Description, infoRes.Book.Duration, infoRes.Book.FileSize, infoRes.Book.PreviewStreamURL)
	Stream(infoRes.Book.PreviewStreamURL)

}
