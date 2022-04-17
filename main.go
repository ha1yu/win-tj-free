package main

import (
	"embed"
	"io/ioutil"
	"log"
)
import _ "embed"

//go:embed tj/FlashCenterInstaller.exe
var embededFiles embed.FS

func main() {
	log.Println(123)
	bytes, err := embededFiles.ReadFile("FlashCenterInstaller.exe")
	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile("/Users/admin/Downloads/xxx.exe", bytes, 0666)
	if err != nil {
		log.Println(err)
		return
	}

}

func main2() {

}
