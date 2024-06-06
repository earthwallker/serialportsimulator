package utils

import (
	"encoding/hex"
	"log"
	"strings"
	_"errors"

	"unicode"
)

/*
*hexString := "0A 03 62 09 24 09 26 09 24 0F D7 0F D7 0F D5 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 13 8A 00 00 00 05 00 00 00 00 00 00 00 09 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 AA AA 00 01 0A 02 00 00 00 00 00 00 00 00 00 00 00 01 00 01 00 00 27 BC"
*/
func Hex2Bytes(hexString string)([]byte,error) {
	

	// 去除空格并分割成十六进制数字
	hexValues := strings.Fields(hexString)
	bytes := make([]byte, len(hexValues))

	for i, hexVal := range hexValues {
		// 将十六进制字符串解析为字节
		byteVal, err := hex.DecodeString(hexVal)
		if err != nil {
			log.Fatalf("hex.DecodeString: %v", err)
			return nil,err
		}

		bytes[i] = byteVal[0]
	}

	// log.Printf("Converted bytes: %v", bytes)
	return bytes,nil
}

/*
*@param input := "00 E6 00 E6 00 E9 00 00 00 00 00 00"
*@return 00e600e600e9000000000000
*/
func ConvertLowerString(input string) string {
	// 去除空格
	input = strings.ReplaceAll(input, " ", "")

	// 将字符串转换为小写
	output := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) {
			return unicode.ToLower(r)
		}
		return r
	}, input)

	return output
}

