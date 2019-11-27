package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Parse(date string) {
	var masser = map[int]string{}
	s := strings.Split(date, "segment")
	z := len(s)
	i := 1
	for i < z {
		masser[i-1] = strings.Split(s[i], ".aac")[0]
		i++
	}
	i = i - 1
	fmt.Println(len(masser))
	fmt.Println(i)

}
func main() {
	file, err := os.Open("parse")
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
	Parse(mass)
}
