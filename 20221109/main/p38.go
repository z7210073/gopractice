package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//查看变量数据类型
	var n1 = 100
	fmt.Printf("n1的类型是 %T \n", n1) //用于格式化输出
	//查看变量占用字节大小
	var n2 int64 = 10000
	fmt.Printf("n2的类型是 %T \n n2占用的字节数 %d ", n2, unsafe.Sizeof(n2))
}
