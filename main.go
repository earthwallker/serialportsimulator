package main

import (
	"serialportsimulator/serial"
	"os"
	"os/signal"
	"syscall"
)


//go:generate bash -c "GOOS=linux GOARCH=arm GOARM=7 go build -o armSimulator"
func main() {
	
	go serial.SerialListen("/dev/ttyS1")
	go serial.SerialListen("/dev/ttyS3")

	sig := make(chan os.Signal, 2)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	<-sig
    destory()
}

func destory() {
	serial.Close("/dev/ttyS1")
	serial.Close("/dev/ttyS3")
}