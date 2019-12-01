package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
)

type GetInfoStruct struct {
	Book    Book `json:"book"`
	Success bool `json:"success"`
}
type Hash struct {
	XClient    string
	XAuthToken string
	XAdToken   string
	UserAgent  string
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

type SocialProfiles struct{}
type CreateNewUser struct {
	Success        bool             `json:"success"`
	Token          string           `json:"token"`
	Username       string           `json:"username"`
	Password       string           `json:"password"`
	SocialProfiles []SocialProfiles `json:"social_profiles"`
	MerchantID     string           `json:"merchant_id"`
}
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func oscheck(err error) {
	if err != nil {
		fmt.Printf("err %s", err)
		os.Exit(1)
	}
}
func GetXAuthToken(XClient, UserAgent string) string {
	url := "https://api.patephone.com/client-api/auth/createNewUser"
	fmt.Println(url)
	client := &http.Client{}
	req, err := http.NewRequest(
		"POST", url, nil,
	)
	req.Header.Add("X-CLIENT-IDENTIFIER", XClient)
	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("X-FEATURES", "F_SHARD,FREE_BOOKS")
	req.Header.Add("Connection", "Keep-Alive")
	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()
	//fmt.Println(resp)
	contents, err := ioutil.ReadAll(resp.Body)
	check(err)
	fmt.Println(string(contents))
	var NewUser CreateNewUser
	err = json.Unmarshal(contents, &NewUser)
	check(err)
	return NewUser.Token
}

type GetAdTokenStruct struct {
	Success bool   `json:"success"`
	Token   string `json:"ad_token"`
}

func GetXAdToken(XClient, UserAgent, XAuthToken string) string {
	url := "https://api.patephone.com/client-api/ad/token"
	fmt.Println(url)
	client := &http.Client{}
	req, err := http.NewRequest(
		"GET", url, nil,
	)
	fmt.Printf("%s %s %s", XClient, UserAgent, XAuthToken)
	req.Header.Add("X-CLIENT-IDENTIFIER", XClient)
	req.Header.Add("User-Agent", UserAgent)
	//	req.Header.Add("X-FEATURES", "F_SHARD,FREE_BOOKS")
	//	req.Header.Add("Connection", "Keep-Alive")
	req.Header.Add("X-AUTH-TOKEN", XAuthToken)
	//	req.Header.Add("Accept-Encoding", "gzip")

	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()
	//fmt.Println("\nToKEN:",resp)
	contents, err := ioutil.ReadAll(resp.Body)
	check(err)
	fmt.Println(string(contents))
	var GetAdToken GetAdTokenStruct
	err = json.Unmarshal(contents, &GetAdToken)
	check(err)
	return GetAdToken.Token
}

type TomlConfig struct {
	XAuthToken string `toml:"XAuthToken"`
	XAdToken   string `toml:"organization"`
}

func Init() (string, string, string, string) {
	//fmt.Printf("Name: '%s', Real: %t, string: %s\n", c.Name, c.Real, world)
	if fileExists("config.toml") {
		fmt.Println("config.toml exists")
		var config TomlConfig
		if _, err := toml.DecodeFile("example.toml", &config); err != nil {
			fmt.Println(err)

		}

		fmt.Printf("XAuthToken: %s XAdToken : %s\n", config.XAuthToken, config.XAdToken)
		XClient := "patephone_unlim_android" //NOT CHANGE !!!!!
		UserAgent := "Patephone Android/8 (XIAOMI Redmi 10 Pro; Android 10)"
		XAuthToken := config.XAuthToken
		XAdToken := config.XAdToken
		return XClient, XAuthToken, XAdToken, UserAgent

	} else {
		fmt.Println("config.toml does not exist ")
		XClient := "patephone_unlim_android" //NOT CHANGE !!!!!
		UserAgent := "Patephone Android/8 (XIAOMI Redmi 10 Pro; Android 10)"
		XAuthToken := GetXAuthToken(XClient, UserAgent)
		XAdToken := GetXAdToken(XClient, UserAgent, XAuthToken)
		f, err := os.Create("config.toml")
		check(err)
		defer f.Close()
		fstring := fmt.Sprintf("XAuthToken = \"%s\" \nXAdToken = \"%s\"\n", XAuthToken, XAdToken)
		_, err = f.WriteString(fstring)
		check(err)
		return XClient, XAuthToken, XAdToken, UserAgent
	}
	//return XClient, XAuthToken, XAdToken, UserAgent
}
func (h Hash) Hello() {
	fmt.Println("Hello ")
}

type Request struct {
	Success bool   `json:"success"`
	Book    []Book `json:"books"`
	Paging  Paging `json:"paging"`
}

func (h Hash) SearchBook(target string) Request {
	url := "https://api.patephone.com/client-api/search/book?search_string=" + target
	client := &http.Client{}
	req, err := http.NewRequest(
		"GET", url, nil,
	)
	req.Header.Add("X-CLIENT-IDENTIFIER", h.XClient)
	req.Header.Add("User-Agent", h.UserAgent)
	req.Header.Add("X-AUTH-TOKEN", h.XAuthToken)
	req.Header.Add("X-AUTH-TOKEN", h.XAdToken)
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

func main() {
	XClient, XAuthToken, XAdToken, UserAgent := Init()
	api := &Hash{
		XClient,
		XAuthToken,
		XAdToken,
		UserAgent,
	}
	api.Hello()
	bookSearch := api.SearchBook("Автостопом")
	fmt.Println(bookSearch)
}
