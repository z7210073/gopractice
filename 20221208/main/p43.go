package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//bool 类型只占一个字节

	var b = false
	fmt.Println(b)

	var a bool = true
	fmt.Println(a)
	fmt.Println("占用空间", unsafe.Sizeof(a))

	s := false
	fmt.Println(s)
}
