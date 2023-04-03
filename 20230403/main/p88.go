package main

import "fmt"

func main() {

	//case 可以是一个常量  有返回值的函数
	var a int = 1
	switch a {
	case ss(0):
		fmt.Println(1)

	}
	//switch 也可以为常量或有返回值的函数
	switch ss(a) {
	case 2:
		fmt.Println(2)

	}

	var n1 int32 = 32
	var n2 int32 = 32
	var n3 int64 = 32
	switch n1 {
	case n2:
		fmt.Println("ok1")

	}
	fmt.Println(n3)
	//类型不一致无法匹配
	//switch n1 {
	//case n3:
	//	fmt.Println("ok1")
	//
	//}

	//case 后面不能有重复的常量

}
func ss(a int) int {
	return a + 1
}
