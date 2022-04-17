package utils

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
)

// ExecShell1 阻塞式的执行外部shell命令的函数,不获取结果,但是会等结果出来之后再返回
func ExecShell1(s string) {
	//函数返回一个*Cmd,用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/sh", "-c", s)

	if runtime.GOOS == "windows" { //如果是windows则修改命令前缀为 cmd /c xxx
		cmd = exec.Command("cmd", "/c", s)
	}

	//Run执行cmd包含的命令,并阻塞直到完成。  这里stdout被取出,cmd.Wait()无法正确获取stdin,stdout,stderr,则阻塞在那了
	if err := cmd.Run(); err != nil {
		log.Println("执行命令出错:", cmd, err)
	}
}

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func UserLocalAppData() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("LOCALAPPDATA")
		return home
	} else {
		return "null"
	}
}

// Write 写入文本
func Write(fileName string, str string) {
	var strByte = []byte(str)
	if !CheckFileExist(fileName) { //写入的文件如果不存在,则创建文件
		file, err := os.Create(fileName) //
		defer file.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}
	err := ioutil.WriteFile(fileName, strByte, 0666)
	if err != nil {
		log.Println(err)
	}
}

// CheckFileExist 检测文件是否存在
func CheckFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
