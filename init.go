package FPatephone

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type SocialProfiles struct{}
type CreateNewUser struct {
	Success        bool             `json:"success"`
	Token          string           `json:"token"`
	Username       string           `json:"username"`
	Password       string           `json:"password"`
	SocialProfiles []SocialProfiles `json:"social_profiles"`
	MerchantID     string           `json:"merchant_id"`
}
type TomlConfig struct {
	XAuthToken string `toml:"XAuthToken"`
	XAdToken   string `toml:"XAdToken"`
	ThreadSize int    `toml:"ThreadSize"`
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

type API struct {
	XClient    string
	XAuthToken string
	XAdToken   string
	UserAgent  string
	ThreadSize int
}

func Api(XClient, XAuthToken, XAdToken, UserAgent string, ThreadSize int) *API {
	return &API{
		XClient:    XClient,
		XAuthToken: XAuthToken,
		XAdToken:   XAdToken,
		UserAgent:  UserAgent,
		ThreadSize: ThreadSize}
}
