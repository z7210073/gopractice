// 该文件所在的包为main包
package main

import "fmt"

// func为关键字，main为主函数，程序的入口。
// 先编译 》 go build hello.go 获得hello.exe文件
// 再运行hello.exe
func main() {
	fmt.Println("hello world")
}
