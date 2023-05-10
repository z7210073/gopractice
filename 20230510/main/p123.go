package main

import "fmt"

// init 函数，每个源文件中都可以有一个init函数，在main函数前运行
func main() {
	fmt.Println("main")
}

func init() {
	fmt.Println("init")
}
