package main

import "fmt"

func main() {
	//制表符 \t
	fmt.Println("tom\tjack")
	fmt.Println("tomjack")
	//换行符 \n
	fmt.Println("hello\nworld")
	fmt.Println("hello world")
	// 输出\
	fmt.Println("\\")

	// 输出双引号
	fmt.Println("I am \"Jason\"")

	//输出回车 \r 从当前行最前面开始输出，覆盖掉以前的内容
	fmt.Println("我z是第一个\r其实我才是第一个")

}
