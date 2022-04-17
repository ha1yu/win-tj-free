package utils

import (
	"log"
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
