package main

import "fmt"

func main() {
	//练习1
	//1.获取int变量num的地址，并显示到终端
	var num int
	fmt.Println(&num)
	fmt.Println(num)
	//2.将num的地址赋值给指针变量ptr，通过ptr修改num的值
	var ptr *int = &num
	*ptr = 10
	fmt.Println(num)
	fmt.Println(ptr)
}
