// test0917 project main.go
package main

import (
	"fmt"
	//	"unsafe"
)

/*const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)*/

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	/*	w1, w2 := swap("hello", "world")
		fmt.Println(w1, w2)
		fmt.Println(split(33))

		//常量可以用len(), cap(), unsafe.Sizeof()常量计算表达式的值。
		//常量表达式中，函数必须是内置函数，否则编译不过
		println(a, b, c)*/

	var a int = 4
	var b int32
	var c float32
	var ptr *int

	/* 运算符实例 */
	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a)
	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b)
	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c)

	/*  & 和 * 运算符实例 */
	ptr = &a /* 'ptr' 包含了 'a' 变量的地址 */
	fmt.Printf("a 的值为  %d\n", a)
	fmt.Printf("*ptr 为 %d\n", *ptr)
}
