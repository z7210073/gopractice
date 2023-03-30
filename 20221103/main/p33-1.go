package main

import "fmt"

func main() {
	//使用变量的三种方式
	// 1. 使用变量声明类型后若不赋值，则使用默认值
	var i int
	fmt.Println(i) // 输出 0
	fmt.Println("---------------------")
	// 2.根据值自行推导类型
	var num = 10
	var str = 's'
	var f = 10.22
	fmt.Println(num)
	fmt.Println(str)
	fmt.Println(f)
	fmt.Println("---------------------")

	// 3. 省略var  使用:= 左侧变量，不应该是已声明过的变量，否则编译无法通过
	name := "jason"
	fmt.Println(name)

}
