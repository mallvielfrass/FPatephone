package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	FPatephone "github.com/mallvielfrass/FPatephone"
)

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
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

type TomlConfig struct {
	XAuthToken string `toml:"XAuthToken"`
	XAdToken   string `toml:"XAdToken"`
	ThreadSize int    `toml:"ThreadSize"`
}

func Init() (string, string, string, string, int) {
	//fmt.Printf("Name: '%s', Real: %t, string: %s\n", c.Name, c.Real, world)
	if fileExists("config.toml") {
		//fmt.Println("config.toml exists")
		var config TomlConfig
		if _, err := toml.DecodeFile("config.toml", &config); err != nil {
			fmt.Println(err)

		}

		fmt.Printf("XAuthToken: %s XAdToken : %s\n", config.XAuthToken, config.XAdToken)
		XClient := "patephone_unlim_android" //NOT CHANGE !!!!!
		UserAgent := "Patephone Android/8 (XIAOMI Redmi 10 Pro; Android 10)"
		XAuthToken := config.XAuthToken
		XAdToken := config.XAdToken
		ThreadSize := config.ThreadSize
		return XClient, XAuthToken, XAdToken, UserAgent, ThreadSize

	} else {
		fmt.Println("config.toml does not exist ")
		XClient := "patephone_unlim_android" //NOT CHANGE !!!!!
		UserAgent := "Patephone Android/8 (XIAOMI Redmi 10 Pro; Android 10)"
		XAuthToken := FPatephone.GetXAuthToken(XClient, UserAgent)
		XAdToken := FPatephone.GetXAdToken(XClient, UserAgent, XAuthToken)
		ThreadSize := 500
		f, err := os.Create("config.toml")
		check(err)
		defer f.Close()
		fstring := fmt.Sprintf("XAuthToken = \"%s\" \nXAdToken = \"%s\"\nThreadSize = %d\n", XAuthToken, XAdToken, ThreadSize)
		_, err = f.WriteString(fstring)
		check(err)
		return XClient, XAuthToken, XAdToken, UserAgent, ThreadSize
	}
	//return XClient, XAuthToken, XAdToken, UserAgent
}
