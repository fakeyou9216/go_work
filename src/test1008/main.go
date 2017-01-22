// test1008 project main.go
package main

import (
	"fmt"
	"strings"
)

func add1(x *int) int {
	*x = *x + 1
	return *x
}

func strUpper(str1 string) string {
	str1 = strings.ToUpper(str1)
	return str1
}

func main() {
	x := 2
	y := add1(&x)
	fmt.Println(x)
	fmt.Println(y)

	strA := "Lion"
	strB := strUpper(strA)
	fmt.Println(strA)
	fmt.Println(strB)
}
