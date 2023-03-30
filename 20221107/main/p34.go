package main

import "fmt"

func main() {
	//变量在该区域的数据值在同一类型范围内可以发生变化
	var i int = 10
	i = 30
	i = 50

	//错误 i=1.2

	fmt.Println(i)

	//变量在同一个作用域里面不能重名
	//错误:'i' redeclared in this block
	//var i int
	//var i string
	// i :=55

	var str string

	fmt.Println(str)
}
