package main

import "fmt"

func main() {
	//指针
	//基本数据类型，变量存的是值，也叫值类型
	//基本数据类型存在内存布局
	var i int = 10
	//i 的地址是什么
	fmt.Println(&i)
	//指针类型，变量存的是一个地址，地址指向的空间才是值
	//ptr与&i相等
	//i的地址是ptr的值，ptr本身就有自己的地址
	var ptr *int = &i
	fmt.Println(ptr)
	fmt.Println(&ptr)
	//获取指针类型的所指向的值需要用*
	//根据指针存的地址指向的对应值
	var num int = *ptr
	fmt.Println(num)
	fmt.Println(*ptr)
}
