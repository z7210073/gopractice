package main

import "fmt"

// 函数也是一种数据类型，他可以赋给变量

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

// 函数可以作为形参作为调用
func myfun(fun func(int, int) int, n1 int, n2 int) int {

	return fun(n1, n2)
}

type myfunType func(int, int) int

func myfun2(fun myfunType, n1 int, n2 int) int {

	return fun(n1, n2)
}

// 通过type 自定义数据类型
// go 支持返回值命名
func calc(n int, v int) (sum int, sub int) {
	sum = n + v
	sub = n - v
	return
}

func main() {

	a := getSum
	//fmt.Printf("%T", a)

	fmt.Println(a(1, 2))
	fmt.Println(myfun(a, 1, 2))
	//go 认为 myInt 和 int 是不同的类型
	type myInt int
	var s myInt
	s = 1
	fmt.Println(s)
	fmt.Println(myfun2(a, 12, 22))

	fmt.Println(calc(10, 2))
}
