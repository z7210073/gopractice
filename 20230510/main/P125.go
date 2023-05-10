package main

import "fmt"

// 闭包是一个函数与其相关的引用环境组合一个整体
//返回一个匿名函数，匿名函数和引用的a 构成一个闭包

func addUpper() func(int) int {
	var a int = 10
	return func(i int) int {
		a = a + i
		return a
	}

}
func main() {
	f := addUpper()   //理解为f 为初始化一个实体，实体中的字段a为同一个a
	fmt.Println(f(1)) //11
	fmt.Println(f(2)) //13
	fmt.Println(f(3)) //16
}
