package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/remeh/sizedwaitgroup"
)

const downloadThreads = 512

var wg = sizedwaitgroup.New(downloadThreads)

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
	XAdToken   string `toml:"XAdToken"`
}

func Init() (string, string, string, string) {
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
	req.Header.Add("X-AD-TOKEN", h.XAdToken)
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
func Parse(id int) map[int]string {
	var masser = map[int]string{}
	path := strconv.Itoa(id) + "/fileList.m3u8"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)
	mass := ""
	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		mass = mass + string(data[:n])
	}
	s := strings.Split(mass, "segment")
	z := len(s)
	i := 1
	for i < z {
		masser[i-1] = "/segment" + strings.Split(s[i], ".aac")[0] + ".aac"
		i++
	}
	i = i - 1
	//fmt.Println(masser)
	//fmt.Printf("%d block\n", i)
	//fmt.Println("---------------------")

	//	fmt.Println("---------------------")
	//	fmt.Println("<-  Download")
	//	Download(masser, id)
	return masser
}
func (h Hash) Stream(id int) {
	//	/stream/hls/preview/13613/fileList.m3u8
	path := strconv.Itoa(id) + "/fileList.m3u8"
	url := "https://n01.cd.ru.patephone.com/stream/hls/ad/" + path
	fmt.Println(url)
	client := &http.Client{}
	req, err := http.NewRequest(
		"GET", url, nil,
	)
	req.Header.Add("X-CLIENT-IDENTIFIER", h.XClient)
	req.Header.Add("User-Agent", h.UserAgent)
	req.Header.Add("X-AUTH-TOKEN", h.XAuthToken)
	req.Header.Add("X-AD-TOKEN", h.XAdToken)
	//	req.Header.Add("Accept-Encoding", "gzip")
	//	req.Header.Add("Connection", "Keep-Alive")
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
	fmt.Println("---------------------")
	fmt.Println("<-  Content")
	cont := string(contents)
	fmt.Println(string(contents))
	fmt.Println("---------------------")
	out, err := os.Create(path)
	check(err)
	defer out.Close()
	_, err = out.WriteString(cont)
	check(err)
	// Write the body to file
	//_, err = io.Copy(out, []byte(cont))
	//check(err)

}
func (h Hash) GetInfo(id int) GetInfoStruct {
	url := "https://api.patephone.com/client-api/books/" + strconv.Itoa(id)
	fmt.Println(url)
	client := &http.Client{}
	req, err := http.NewRequest(
		"GET", url, nil,
	)
	req.Header.Add("X-CLIENT-IDENTIFIER", h.XClient)
	req.Header.Add("User-Agent", h.UserAgent)
	req.Header.Add("X-AUTH-TOKEN", h.XAuthToken)
	req.Header.Add("X-AD-TOKEN", h.XAdToken)
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
	//Stream(infoRes.Book.ID) // MUST BE OPTIMIZED SEND Id
	return infoRes
}
func ChlkDir(id int) {
	path := strconv.Itoa(id)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0777)
	}
}
func (h Hash) DownloadFile(path string) error {
	defer wg.Done()

	// Get the data
	url := "https://n01.cd.ru.patephone.com/stream/hls/ad/" + path
	fmt.Println(url)
	client := &http.Client{}
	req, err := http.NewRequest(
		"GET", url, nil,
	)
	req.Header.Add("X-CLIENT-IDENTIFIER", h.XClient)
	req.Header.Add("User-Agent", h.UserAgent)
	req.Header.Add("X-AUTH-TOKEN", h.XAuthToken)
	req.Header.Add("X-AD-TOKEN", h.XAdToken)
	req.Header.Add("Accept-Encoding", "identity")
	req.Header.Add("Connection", "Keep-Alive")
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
func (h Hash) Download(dict map[int]string, id int) {
	fmt.Println(dict)
	n := len(dict)
	for i := 0; i < n; i++ {
		wg.Add()
		path := strconv.Itoa(id) + dict[i]
		go func() {
			if err := h.DownloadFile(path); err != nil {
				panic(err)
			}
		}()
	}
	wg.Wait()
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
	bookSearch := api.SearchBook("time")
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
	fmt.Println("---------------------")
	api.Download(parseList, info.Book.ID)

}
