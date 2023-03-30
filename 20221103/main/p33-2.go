package main

import "fmt"

// 声明全局变量
var s1 = 1
var s2 = 2
var name3 = "jason"

// 一次性声明全局变量
var (
	s11   = 11
	s22   = 22
	name4 = "jason2"
)

func main() {
	//同类型多变量声明

	var n1, n2, n3 int
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	//不同类型多变量声明
	var num, str = 100, "jason"

	fmt.Println(num)
	fmt.Println(str)

	//使用类型推导
	//不同类型多变量声明
	num2, str2 := 1002, "jason2"

	fmt.Println(num2)
	fmt.Println(str2)
	fmt.Println("---------------")

	fmt.Println(s1, s2, name3)
	fmt.Println("---------------")
	fmt.Println(s11, s22, name4)
}
