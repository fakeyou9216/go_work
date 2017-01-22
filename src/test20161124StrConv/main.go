// test20161124StrConv project main.go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	rs := []rune("golang中文unicode编码")
	json := ""
	html := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			json += string(r)
			html += string(r)
		} else {
			json += "\\u" + strconv.FormatInt(int64(rint), 16) // json
			html += "&#" + strconv.Itoa(int(r)) + ";"          // 网页
		}
	}
	fmt.Printf("JSON: %s\n", json)
	fmt.Printf("HTML: %s\n", html)
}
