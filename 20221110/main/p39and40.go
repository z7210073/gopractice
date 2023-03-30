package main

import "fmt"

func main() {
	//浮点类型

	//类似单精度浮点型
	var i float32 = 1.22
	fmt.Println(i)

	//类似双精度浮点型 java的double
	var k float64 = 1.66666666
	fmt.Println(k)

	//浮点数可能会造成精度损失
	var n1 float32 = -123.0000901
	var n2 float64 = -123.0000901

	fmt.Println("n1=", n1, "n2=", n2)
	//输出结果 ： n1= -123.00009 n2= -123.0000901

	//go 默认声明浮点型为float64
	n5 := 1.22

	fmt.Printf("n5类型为%T \n", n5)
	//输出 .123 等价0.123
	n6 := .123
	fmt.Println(n6)

}
