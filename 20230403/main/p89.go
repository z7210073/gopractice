package main

import "fmt"

func main() {
	//switch 可以不带 表达式
	//switch 后面也可以声明一个变量
	var age int = 10
	switch {
	case age == 10:
		fmt.Println("10")

	}

	//case可以对范围进行判断
	switch {
	case age > 9:
		fmt.Println(">10")
	}
	var money int = 1000
	//switch 穿透 fallthrough,即匹配第一个后仍然继续匹配
	var money2 int = 1000
	switch money {
	case 1000:
		fmt.Println("10000")
		fallthrough
	case money2:
		fmt.Println("ok2")
	}
}
