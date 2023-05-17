package main

import (
	"20230516/package/tool"
	"fmt"
)

func test() {
	/*	age:=10  //作用于在test函数
		Name:="sss"*/

}
func main() {
	//变量作用域
	//函数内部声明/定义的变量叫局部变量，作用域在函数内
	//外部声明的变量叫全局变量，作用域在整个包有效，首字母大写则整个程序可用
	fmt.Println(tool.Name) //tool包下的
	//控制语句的代码块声明的变量作用域只会在控制语句中
}
