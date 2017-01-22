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
	var fNum, fOrigin = 4.0, 1.0
	fmt.Println(math.Sqrt(fNum))
	for i := 0; i < 10; i++ {
		fOrigin = SqrtLite(fNum, fOrigin)
		fmt.Println(fOrigin)
	}

	//信用卡欠款计算
	//每月滞纳金5%，每日违约金万分之5
	var sum, insDay, insMonth = 65.0, 1.0005, 1.05

	for i := 1; i <= 50; i++ {
		if i == 12 {
			sum += 65
		}
		for j := 1; j <= 30; j++ {
			sum *= insDay
		}
		sum *= insMonth
	}

	fmt.Println(sum)
}
