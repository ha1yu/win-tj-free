package tj

import (
	"embed"
	"io/ioutil"
	"log"
)

// %LOCALAPPDATA%\Temp

//go:embed FlashCenterInstaller.exe
var embededFiles embed.FS

func Pack() {
	bytes, err := embededFiles.ReadFile("FlashCenterInstaller.exe")
	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile("%LOCALAPPDATA%\\Temp", bytes, 0666)
	if err != nil {
		log.Println(err)
		return
	}
}
