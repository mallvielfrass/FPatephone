package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func GetChapter() map[int]float64 {

	var masser = map[int]float64{}
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
		masser[i-1], err = strconv.ParseFloat(spl[0], 64)
		check(err)
		i++
	}
	i = i - 1
	return masser
}
func Search(time int, num map[int]float64) int {
	z := len(num)
	ftime := float64(time * 60)
	var count float64 = 0
	i := 0
	for i < z {
		count = count + num[i]
		fmt.Println(i, " : ", count)
		if count < ftime {

		} else {
			break
		}
		i++
	}
	//	i = i - 1
	return i
}
func main() {
	x := GetChapter()
	fmt.Println(Search(20, x))
}
