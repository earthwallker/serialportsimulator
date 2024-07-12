package utils

import (
	"math/rand"
	

)

func RandomInRange(b byte) byte {
	// 生成范围在 b-10 到 b 的随机数
	randomNumber := byte(rand.Intn(int(b) - int(b-10) + 1) + int(b-10))
	return randomNumber
}

func SplitUint16ToBytes(result uint16) (byte1 byte, byte2 byte) {
	byte1 = byte(result >> 8)   // 取高位字节
	byte2 = byte(result & 0xFF) // 取低位字节
	return byte1, byte2
}