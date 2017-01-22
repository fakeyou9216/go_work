// test1107TcpSocketClient project main.go
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	os.Args = []string{"MyPc", "计算机localhost:7777"}
	if len(os.Args) <= 1 {
		fmt.Println("usage: go run client3.go YOUR_CONTENT")
		return
	}
	log.Println("begin dial...")
	conn, err := net.Dial("tcp", "192.168.1.132:8888")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	defer conn.Close()
	log.Println("dial ok")

	time.Sleep(time.Second * 2)
	data := os.Args[1]
	conn.Write([]byte(data))
}
