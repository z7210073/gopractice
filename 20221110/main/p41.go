package main

import "fmt"

func main() {
	//go 没有专门的字符类型
	//go的字符串是由字节组成的

	// 保存小写字母a

	var c1 byte = 'a'
	fmt.Println(c1)
	var c2 byte = '0'
	fmt.Println(c2)
	//当我们直接输出byte ascII码表的值

	//如果输出对应的字符
	fmt.Printf("%c \n", c1)

	var cn rune = '北'
	fmt.Printf("%c", cn)
	fmt.Printf("%d", cn)
}
