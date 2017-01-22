// test1103IPConvert project main.go
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	os.Args = []string{"MyPc", "127.0.0.1"}
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)
}
