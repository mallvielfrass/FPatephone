package FPatephone

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func (api *API) Stream(id int) string {
	//	/stream/hls/preview/13613/fileList.m3u8
	path := strconv.Itoa(id) + "/fileList.m3u8"
	url := "https://n01.cd.ru.patephone.com/stream/hls/ad/" + path
	fmt.Println(url)
	client := &http.Client{}
	req, err := http.NewRequest(
		"GET", url, nil,
	)
	req.Header.Add("X-CLIENT-IDENTIFIER", api.XClient)
	req.Header.Add("User-Agent", api.UserAgent)
	req.Header.Add("X-AUTH-TOKEN", api.XAuthToken)
	req.Header.Add("X-AD-TOKEN", api.XAdToken)
	//	req.Header.Add("Accept-Encoding", "gzip")
	//	req.Header.Add("Connection", "Keep-Alive")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error", err)
		return "err"
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err %s", err)
		os.Exit(1)
	}
	cont := string(contents)
	return cont

}
