package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"serialportsimulator/model"
	"serialportsimulator/mymodbuscrc"
	"serialportsimulator/utils"
	_ "time"

	// "github.com/jacobsa/go-serial/serial"
	"github.com/goburrow/serial"
	"github.com/logrusorgru/aurora"
	
)



var(
  	deviceProtocols []model.DeviceProtocol
)

func init(){
	deviceProtocols=model.Protocols
}

//go:generate bash -c "GOOS=linux GOARCH=arm GOARM=7 go build -o armSimulator"
func main() {
	// 打开串口
	port, err := serial.Open(&serial.Config{
		Address: "/dev/ttyS1",
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

	buf := make([]byte, 100)
	
	// utils.Hex2Bytes(hexString)

	for {
		n, err := port.Read(buf)
		if err != nil {
			log.Fatalf("port.Read: %v", err)
		}

		if n > 0 {
			isExistDev:=false
			// slaveID+Fun+data+CRC:= buf[:n]
			// data := buf[:n-2]
			slaveId_fun:= buf[:2]
			
            var datas []byte
			fun_reg_count:= buf[1:6]
			hexString := hex.EncodeToString(fun_reg_count)
			devName:="无设备："
            for i:=0;i<len(deviceProtocols);i++ {
				command:=utils.ConvertLowerString(deviceProtocols[i].Cmd)
				if hexString == command {
					isExistDev=true
					devName =deviceProtocols[i].Tag
					datas,err=utils.Hex2Bytes(deviceProtocols[i].Data)
					if err != nil {
						log.Fatal(err)
					}
				}
			}

			packets := slaveId_fun
			if isExistDev {
				packets = append(packets, byte(len(datas)))
				packets = append(packets, datas...)
				crc1,crc2:=mymodbuscrc.ModbusCrc(packets)
			
				packets = append(packets, crc1)
				packets = append(packets, crc2)
				// time.Sleep(time.Millisecond * 100) // 等待一段时间后继续读取，可以根据实际情况调整
			}else{
				msg:=fmt.Sprintf("Received hex data: %s\n", hexString)
				fmt.Println(aurora.Red(msg))
			}

			// 向串口写入数据
			_, err := port.Write(packets)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(aurora.Cyan(devName),aurora.Blue(hexString),aurora.Green(packets[:3]))
		}
	}


}