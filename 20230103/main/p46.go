package main

import "fmt"

func main() {
	//输出基本数据类型的默认值  在go语言中默认值也叫零 值
	var num int      //0
	var num2 float64 //0
	var str string   //""
	var s bool       // false

	fmt.Println(num)
	fmt.Println(num2)
	fmt.Println(str)
	fmt.Println(s)

}
