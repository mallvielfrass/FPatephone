package main

import (
	"fmt"
	"os"
	"strconv"

	FPatephone "github.com/mallvielfrass/FPatephone"
)

func main() {
	//FPatephone
	XClient, XAuthToken, XAdToken, UserAgent, ThreadSize := Init()
	api := FPatephone.Api(XClient, XAuthToken, XAdToken, UserAgent, ThreadSize)
	argsv := os.Args[1:]
	searchWord := ""
	for wrld := range argsv {

		sp := ""
		if wrld == 0 {

		} else {
			sp = " "
		}
		searchWord = searchWord + sp + argsv[wrld]
	}
	fmt.Println(searchWord)
	if searchWord != "" {

	} else {
		searchWord = "сказка"
	}
	bookSearch := api.SearchBook(searchWord)
	n := bookSearch.Paging.Count
	if n > 0 {
		for i := 0; i < n; i++ {
			fmt.Println(i, "\b.) ", bookSearch.Book[i].ID, " : ", bookSearch.Book[i].Title) //9411
		}
	} else {
		fmt.Println("book not found")
	}
	//fmt.Println("debug")
	idNum := Select()
	fmt.Println("you selected:", idNum, "\b.)", bookSearch.Book[idNum].Title)
	fmt.Println("---------------------")
	fmt.Println("---------------------")
	info := api.GetInfo(bookSearch.Book[idNum].ID)
	fmt.Printf(">>%d \n>>%s \n>>%s \n>>%s \n>>%s \n>>%d \n>>%d \n>>%s \n", info.Book.ID, info.Book.Title, info.Book.Authors[1].FirstName, info.Book.Authors[1].LastName, info.Book.Description, info.Book.Duration, info.Book.FileSize, info.Book.PreviewStreamURL)
	fmt.Println("<-  ChlkDir")
	ChlkDir(info.Book.ID)
	fmt.Println("---------------------")
	//StreamDownload(info.Book.ID, api)
	//	func StreamDownload(id int, api *API) {
	stream := api.Stream(info.Book.ID)
	patch := strconv.Itoa(info.Book.ID) + "/fileList.m3u8"
	out, err := os.Create(patch)
	check(err)
	defer out.Close()
	_, err = out.WriteString(stream)
	check(err)
	infoB := api.GetChapter(info.Book.ID)
	for i := 0; i < len(infoB.Chapters); i++ {
		fmt.Println(infoB.Chapters[i].Name)
	}
	//}
}
