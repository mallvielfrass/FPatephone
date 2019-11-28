package main

import (
	"encoding/json"
	"fmt"
	"io"
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

func DownloadFile(path string) error {

	// Get the data
	url := "https://n01.cd.ru.patephone.com/stream/hls/ad/" + path
	fmt.Println(url)
	client := &http.Client{}
	req, err := http.NewRequest(
		"GET", url, nil,
	)
	req.Header.Add("X-CLIENT-IDENTIFIER", "patephone_unlim_android")
	req.Header.Add("Accept-Encoding", "identity")
	req.Header.Add("X-AUTH-TOKEN", "hl2LqwCUdcw3RXMT7Rr34NXuULt2jXD3ymnYXqLN65HbeV4zjMCMjcgDFbysKdeZ")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error", err)
		return err
	}
	defer resp.Body.Close()
	fmt.Println(resp)
	// Create the file

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
func Download(dict map[int]string, id int) {
	fmt.Println(dict)
	n := len(dict)
	path := ""
	for i := 0; i < n; i++ {
		path = strconv.Itoa(id) + "/segment" + dict[i] + ".aac"
		if err := DownloadFile(path); err != nil {
			panic(err)
		}
		//fmt.Println(path)
		//	url := "https://n01.cd.ru.patephone.com/stream/hls/ad/" + path
		//	fmt.Println(url)
		//	client := &http.Client{}
		//	req, err := http.NewRequest(
		//		"GET", url, nil,
		//	)
		//	req.Header.Add("X-CLIENT-IDENTIFIER", "patephone_unlim_android")
		//
		//	resp, err := client.Do(req)
		//	if err != nil {
		//		fmt.Println("error", err)
		//		return
		//	}
		//	defer resp.Body.Close()
		//	out, _ := os.Create(path) //ALARM Fix it

		//	defer out.Close()
		// Write the body to file
		//	fmt.Println("---------------------")
		//	fmt.Println("<-  resp.Body")
		//	fmt.Println(resp.Body)
		//	_, err = io.Copy(out, resp.Body)
		//	return err //ALARM Fix it

	}
}
func Parse(date string, id int) {
	var masser = map[int]string{}
	s := strings.Split(date, "segment")
	z := len(s)
	i := 1
	for i < z {
		masser[i-1] = strings.Split(s[i], ".aac")[0]
		i++
	}
	i = i - 1
	//fmt.Println(masser)
	fmt.Printf("%d block\n", i)
	fmt.Println("---------------------")
	fmt.Println("<-  ChlkDir")
	ChlkDir(id)
	fmt.Println("---------------------")
	fmt.Println("<-  Download")
	Download(masser, id)
}
func ChlkDir(id int) {
	path := strconv.Itoa(id)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0777)
	}
}
func Stream(id int) { // MUST BE OPTIMIZED generate link from id
	//	/stream/hls/preview/13613/fileList.m3u8

	url := "https://n01.cd.ru.patephone.com/stream/hls/preview/" + strconv.Itoa(id) + "/fileList.m3u8"
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
	//	fmt.Println(string(contents))
	fmt.Println("---------------------")
	fmt.Println("<-  Parse")
	Parse(string(contents), id)

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
	Stream(infoRes.Book.ID) // MUST BE OPTIMIZED SEND Id

}
