package main

import (
	"fmt"
	"math"
)

func SqrtLite(x, y float64) float64 {
	y = y - (y*y-x)/(2*y)
	return y
}

func main() {

	//牛顿逐次逼近法求平方根
	var fNum, fOrigin = 96.0, 1.0
	fmt.Println(math.Sqrt(fNum))
	for i := 0; i < 10; i++ {
		fOrigin = SqrtLite(fNum, fOrigin)
		fmt.Println(fOrigin)
	}
}
