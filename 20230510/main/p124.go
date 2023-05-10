package main

import "fmt"

//全局匿名函数

var (
	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

// 匿名函数
// 某个函数只希望使用一次
func main() {
	//变量a为匿名函数
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	//求两个数的和
	rustl := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println(rustl)

	fmt.Println(a(20, 10))

	fmt.Println(Fun1(2, 3))
}

//
//全局匿名函数
