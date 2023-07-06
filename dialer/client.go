package main

import (
	"time"

	"github.com/cloudwego/netpoll"
)

func main() {
	network, address, timeout := "tcp", "127.0.0.1:8080", 50*time.Hour

	// use dialer
	dialer := netpoll.NewDialer()
	conn, _ := dialer.DialConnection(network, address, timeout)

	// write & send message
	writer := conn.Writer()
	_, err := writer.WriteString("hello world")
	if err != nil {
		return
	}
	err = writer.Flush()
	if err != nil {
		return
	}
}
