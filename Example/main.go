package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type TomlConfig struct {
	XAuthToken string `toml:"XAuthToken"`
	XAdToken   string `toml:"XAdToken"`
	ThreadSize int    `toml:"ThreadSize"`
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func Init() (string, string, string, string, int) {
	//fmt.Printf("Name: '%s', Real: %t, string: %s\n", c.Name, c.Real, world)
	if fileExists("config.toml") {
		fmt.Println("config.toml exists")
		var config TomlConfig
		if _, err := toml.DecodeFile("config.toml", &config); err != nil {
			fmt.Println(err)

		}

		fmt.Printf("XAuthToken: %s XAdToken : %s\n", config.XAuthToken, config.XAdToken)
		XClient := "patephone_unlim_android" //NOT CHANGE !!!!!
		UserAgent := "Patephone Android/8 (XIAOMI Redmi 10 Pro; Android 10)"
		XAuthToken := config.XAuthToken
		XAdToken := config.XAdToken
		ThreadSize := config.ThreadSize
		return XClient, XAuthToken, XAdToken, UserAgent, ThreadSize

	} else {
		fmt.Println("config.toml does not exist ")
		XClient := "patephone_unlim_android" //NOT CHANGE !!!!!
		UserAgent := "Patephone Android/8 (XIAOMI Redmi 10 Pro; Android 10)"
		XAuthToken := GetXAuthToken(XClient, UserAgent)
		XAdToken := GetXAdToken(XClient, UserAgent, XAuthToken)
		ThreadSize := 500
		f, err := os.Create("config.toml")
		check(err)
		defer f.Close()
		fstring := fmt.Sprintf("XAuthToken = \"%s\" \nXAdToken = \"%s\"\nThreadSize = %d\n", XAuthToken, XAdToken, ThreadSize)
		_, err = f.WriteString(fstring)
		check(err)
		return XClient, XAuthToken, XAdToken, UserAgent, ThreadSize
	}
	//return XClient, XAuthToken, XAdToken, UserAgent
}
func main() {
	XClient, XAuthToken, XAdToken, UserAgent, ThreadSize := Init()
	api := &Hash{
		XClient,
		XAuthToken,
		XAdToken,
		UserAgent,
		ThreadSize,
	}
	api.Hello()
	bookSearch := api.SearchBook("сказка")
	n := bookSearch.Paging.Count
	if n > 0 {
		for i := 0; i < n; i++ {
			fmt.Println(i, "\b.) ", bookSearch.Book[i].ID, " : ", bookSearch.Book[i].Title) //9411
		}
	} else {
		fmt.Println("book not found")
	}
	idNum := Select()
	fmt.Println("you selected:", idNum, "\b.)", bookSearch.Book[idNum].Title)
	fmt.Println("---------------------")
	info := api.GetInfo(bookSearch.Book[idNum].ID)
	fmt.Printf(">>%d \n>>%s \n>>%s \n>>%s \n>>%s \n>>%d \n>>%d \n>>%s \n", info.Book.ID, info.Book.Title, info.Book.Authors[1].FirstName, info.Book.Authors[1].LastName, info.Book.Description, info.Book.Duration, info.Book.FileSize, info.Book.PreviewStreamURL)
	fmt.Println("<-  ChlkDir")
	ChlkDir(info.Book.ID)
	fmt.Println("---------------------")
	api.Stream(info.Book.ID)
	parseList := Parse(info.Book.ID)
	fmt.Println("parseList:\n", parseList)

	api.Download(parseList, info.Book.ID)
	fmt.Println("---------------------")
	fmt.Println("finish")

}
