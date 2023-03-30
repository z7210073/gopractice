package main

import (
	"fmt"
	"strconv"
)

func main() {
	//string 转基本数据类型，不能有效转换的时候则转换成默认值
	var s string = "hello"
	var i int64
	i, _ = strconv.ParseInt(s, 10, 64)

	fmt.Println(i)
}
