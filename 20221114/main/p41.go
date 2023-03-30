package main

import "fmt"

func main() {
	//go语言是以UTF-8编码

	//中文与编码之间的转换
	var num rune = 20013
	fmt.Printf("%c \n", num)

	var cn rune = '中'
	fmt.Printf("%d\n", cn)

}
