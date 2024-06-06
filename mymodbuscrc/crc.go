package mymodbuscrc

import (
	"fmt"
	"hash/crc32"
	"hash/crc64"
)

/*
*@param data []byte{0x01, 0x02, 0x03, 0x04, 0x05}
*/
func Crc32(data []byte) {
	// 计算CRC32校验值
	crc32q := crc32.MakeTable(0xD5828281)
	crc := crc32.Checksum(data, crc32q)

	fmt.Printf("原始数据: %v\n", data)
	fmt.Printf("CRC32校验值: %x\n", crc)
}

func Crc16ISO(data []byte) {
	// 计算CRC16校验值
	crc16Table := crc64.MakeTable(crc64.ISO)
	crc := crc64.Checksum(data, crc16Table)

	fmt.Printf("原始数据: %v\n", data)
	fmt.Printf("CRC16校验值: %x\n", crc)
}

/*
*@param data []byte{0x0A, 0x03, 0x00, 0xE0, 0x00, 0x06}
*/
func Crc16ECMA(data []byte) {
	// 计算CRC16校验值
	crc16Table := crc64.MakeTable(crc64.ECMA)
	crc := crc64.Checksum(data, crc16Table)

	// 分离高位字节和低位字节
	crcHigh := byte(crc >> 8)
	crcLow := byte(crc)

	fmt.Printf("原始数据: %v\n", data)
	fmt.Printf("计算得到的CRC16校验值: %02X %02X\n", crcHigh, crcLow)
}


// 计算Modbus CRC校验值
func calculateCRC(data []byte) uint16 {
	var crc uint16 = 0xFFFF

	for _, b := range data {
		crc ^= uint16(b)
		for i := 0; i < 8; i++ {
			if crc&0x0001 != 0 {
				crc >>= 1
				crc ^= 0xA001
			} else {
				crc >>= 1
			}
		}
	}

	return crc
}

/*0A 03 00 E0 00 06 C5 45 
*@param data := []byte{0x0A, 0x03, 0x00, 0xE0, 0x00, 0x06}
*@return C5 45 
*/
func ModbusCrc(data[]byte) (byte,byte){

	// 计算Modbus CRC校验值
	crc := calculateCRC(data)

	// 分离高位字节和低位字节
	crcLow := byte(crc & 0xFF)
	crcHigh := byte((crc >> 8) & 0xFF)

	// fmt.Printf("原始数据: %v\n", data)
	// fmt.Printf("计算得到的Modbus CRC校验值: %02X %02X\n", crcLow,crcHigh)
	return crcLow,crcHigh
}