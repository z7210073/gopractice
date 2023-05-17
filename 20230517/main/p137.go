package main

import (
	"fmt"
	"strings"
)

func main() {
	//1.去掉左右两边空格
	fmt.Println(strings.TrimSpace("  a ssdsa  dsa  "))
	//指定去掉两边的
	fmt.Println(strings.Trim("! hello !", " hoe!"))
	//去掉左边
	fmt.Println(strings.TrimLeft("! hello !", " hoe!"))
	//去掉右边
	fmt.Println(strings.TrimRight("! hello !", " hoe!"))
	//判断是否以某个开头
	fmt.Println(strings.HasPrefix("qwefsad", "qwe"))
	//判断是否以某个结尾

	fmt.Println(strings.HasSuffix("feweqwq", "wq"))

}
