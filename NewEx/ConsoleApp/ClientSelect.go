package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
