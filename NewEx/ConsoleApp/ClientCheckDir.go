package main

import (
	"os"
	"strconv"
)

func ChlkDir(id int) {
	path := strconv.Itoa(id)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0777)
	}
}
