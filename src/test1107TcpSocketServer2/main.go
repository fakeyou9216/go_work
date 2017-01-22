// test1107TcpSocketServer2 project main.go
package main

import (
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		// read from the connection
		time.Sleep(1 * time.Second)
		var buf = make([]byte, 30)
		log.Println("start to read from conn")
		n, err := c.Read(buf)
		if err != nil {
			log.Println("conn read error:", err)
			return
		}
		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
	}
}

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Println("listen error:", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("accept error:", err)
			break
		}
		// start a new goroutine to handle
		// the new connection.
		//显示服务器的网络名和地址
		log.Println("accept a new connection %s", l.Addr().String())
		log.Println("RemoteAddr is %s", c.RemoteAddr())
		go handleConn(c)
	}
}
