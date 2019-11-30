package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

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

func check(e error) {
	if e != nil {
		panic(e)
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
	fmt.Printf("%s %s %s",XClient, UserAgent, XAuthToken)
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
func Init() (string, string, string, string) {
	//fmt.Printf("Name: '%s', Real: %t, string: %s\n", c.Name, c.Real, world)
	if fileExists("config.toml") {
		fmt.Println("config.toml exists")
		return "XClient", "XAuthToken", "XAdToken", "UserAgent"
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
func main() {
	XClient, XAuthToken, XAdToken, UserAgent := Init()
	api := &Hash{
		XClient,
		XAuthToken,
		XAdToken,
		UserAgent,
	}
	api.Hello()
}
