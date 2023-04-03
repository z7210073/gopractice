package main

import "fmt"

func main() {

	//与java不同，匹配项case结束后不用加break
	a := 4
	switch a {
	case 1:
		println(1)
	case 2, 3: //case 可以多个匹配值
		println(2)
	default: //都不匹配执行default 语句  （不是必须出现）
		println(321)
	}
	var d byte
	fmt.Println("输入一个字母a~e")
	fmt.Scanf("%c", &d)
	switch d {
	case 'a':
		fmt.Println("zhouyi")
	case 'b':
		fmt.Println("zhouer")

	}
}
