package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func GetChapter() map[int]string {

	var masser = map[int]string{}
	//	path := strconv.Itoa(id) + "/fileList.m3u8"

	file, err := os.Open("fileList.m3u8")
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
	s := strings.Split(mass, "#EXTINF:")
	z := len(s)
	i := 1

	for i < z {
		spl := strings.Split(s[i], ",")
		masser[i-1] = spl[0] + ":" + strings.Split(spl[1], "\n")[1]
		i++
	}
	i = i - 1
	return masser
}
func main() {
	fmt.Println(GetChapter())
}
