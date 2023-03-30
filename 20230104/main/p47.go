package main

import "fmt"

func main() {
	//go 语言中不数据类型不支持自动转换，需要显式转换（强制转换）

	//演示资本数据类型的转换
	var i int = 100
	//将i转换成float ,被转换的是值，i的数据类型没有变化
	var n float64 = float64(i)
	fmt.Println(n)

	//精度转换
	var i2 int32 = 1
	var i3 int8 = int8(i2)
	fmt.Println(i3)
}
