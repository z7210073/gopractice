package main

import "fmt"

func main() {
	//整数类型使用
	var i int = 1
	fmt.Println("i=", i)

	//测试int8范围(有符号）
	var j int8 = 127

	//	var j int8 = -129
	//The value of '-129' (type untyped int) cannot be represented by the type int8
	fmt.Println(j)

	//测试uint8 范围（无符号）
	var k uint8 = 255
	fmt.Println(k)
	//最大为255
	//'256' (type untyped int) cannot be represented by the type uint8
	//var k uint8=-1
	//最小为0
	// The value of '-1' (type untyped int) cannot be represented by the type uint8

	//int ,uint,rune,byte
	var a int = 8905
	fmt.Println(a)
	var b uint = 9000
	fmt.Println(b)

	var c byte = 255 //0~255
	fmt.Println(c)

}
