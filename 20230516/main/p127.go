package main

import "fmt"

// defer 用法 在函数执行完毕时候会执行该语句
func sum(n1 int, n2 int) int {
	//执行到defer时，会将defer押入栈中
	//defer 执行顺序为先入后出
	defer fmt.Println("ok1")
	defer fmt.Println("ok2")
	res := n1 + n2
	fmt.Println("ok4")
	fmt.Println(res)
	defer fmt.Println("ok3")
	return res

}

func main() {
	//打印顺序 ok4   3  ok3 ok2 ok1  hello
	sum(1, 2)
	fmt.Println("hello ")
}
