package main

import "fmt"

// 引用改变
// 函数传递一般是值拷贝
func s(n *int) {

	*n = 10

	//return  n
}

func main() {

	n := 11
	fmt.Println(n)
	s(&n)
	fmt.Println(n)
}
