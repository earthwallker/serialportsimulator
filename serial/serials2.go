package serial

import (
	_"encoding/hex"
	"fmt"
	"log"
	"serialportsimulator/model"
	"serialportsimulator/mymodbuscrc"
	_"serialportsimulator/utils"
	"time"
	"math"

	// "github.com/jacobsa/go-serial/serial"
	"github.com/goburrow/serial"
	"github.com/logrusorgru/aurora"
	
)
var(
	deviceProtocols []model.DeviceProtocol
	serialPort map[string]serial.Port
)

func init(){
  deviceProtocols=model.Protocols
  serialPort=make(map[string]serial.Port)
}

//dev/ttyS1
//dev/ttyS3
func SerialListen(comName string) {
	// 打开串口
	port, err := serial.Open(&serial.Config{
		Address: comName,
		// Device path (/dev/ttyS0)
		// Baud rate (default 19200)
		BaudRate :9600,
		// Data bits: 5, 6, 7 or 8 (default 8)
		DataBits :8,
		// Stop bits: 1 or 2 (default 1)
		StopBits :1,
		// Parity: N - None, E - Even, O - Odd (default E)
		// (The use of no parity requires 2 stop bits.)
		Parity :"N",
		// Read (Write) timeout.
	})

	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

    fmt.Println("打开串口：",comName)
	serialPort[comName]=port


	// utils.Hex2Bytes(hexString)

	for {
		buf := make([]byte, 100)
		packets := make([]byte, 5)

		n, err := port.Read(buf)
		if err != nil {
			log.Fatalf("port.Read: %v", err)
		}

		if n > 0 {
	
				// slaveID+Fun+data+CRC:= buf[:n]
				// data := buf[:n-2]

				// 解析Modbus RTU消息
				// Modbus RTU消息格式通常是：Slave ID | Function Code | Data | CRC Checksum
				// 在这个例子中，我们假设Data部分包含两个字节的寄存器地址和两个字节的寄存器数量
				fmt.Println(aurora.Red(buf[:n-2]),aurora.Cyan(buf[n-2:n]))
				if len(buf) < 5 {
					log.Fatalf("invalid data length: expected at least 5 bytes, got %d", len(buf))
				}

				slaveID := buf[0]
				functionCode := buf[1]
				registerAddress := uint16(buf[2])<<8 | uint16(buf[3])
				registerCount := uint16(buf[4])<<8 | uint16(buf[5])
				var count int
				if functionCode==3 {
					count= (int)(registerCount*2)
				}

				if functionCode==2 {
					byteCount := int(math.Ceil(float64(registerCount) / 8.0))
					count = byteCount
				}
				
				data := make([]byte, count)

				fmt.Printf("Slave ID: %d\n", slaveID)
				fmt.Printf("Function Code: %d\n", functionCode)
				fmt.Printf("Register Address: %d\n", registerAddress)
				fmt.Printf("Register Count: %d\n", registerCount)

				packets = append(packets, slaveID)
				packets = append(packets, functionCode)
				packets = append(packets, byte(registerCount*2))
				
				packets = append(packets, data...)
				crc1,crc2:=mymodbuscrc.ModbusCrc(packets)
			
				packets = append(packets, crc1)
				packets = append(packets, crc2)

				time.Sleep(time.Millisecond * 100) // 等待一段时间后继续读取，可以根据实际情况调整

				// 向串口写入数据
				_, err := port.Write(packets)
			
				if err != nil {
					log.Fatal(err)
				}


		}
	}








}
    
	



func Close(comName string) {
	if serial,ok :=serialPort[comName];ok{
       serial.Close()
	}
}