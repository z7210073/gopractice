package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 字符串函数
func main() {
	//字符串长度 中文长度 一个汉字占3个字节
	var str string = "hello"
	fmt.Println(len(str))
	//字符串遍历 含中文问题
	var str2 string = "hello北京"
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("str=%c ", r[i])
	}
	fmt.Println("")
	//字符串转整数
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(n)
	}
	//整数转字符穿
	str4 := strconv.Itoa(123)
	fmt.Printf("%T", str4)

	//字符串转 byte
	var bytes = []byte("hello")
	fmt.Printf("bytres=%v \n", bytes)

	//byte 转换成字符串
	str = string([]byte{97, 98, 99, 100})
	fmt.Println(str)
	//进制转换
	str = strconv.FormatInt(123, 16)
	fmt.Println(str)
	//包含字符串
	var con string = "ssIamsss"
	fmt.Println(strings.Contains(con, "Iam"))
	//计算几个子串
	fmt.Println(strings.Count(con, "s"))
	//区分大小写字符串比较
	a := "sssSSS我"
	c := "sssSSS我"
	b := "ssSSSS我"
	fmt.Println(a == c)
	fmt.Println(a == b)
	//不区分大小写比较
	fmt.Println(strings.EqualFold(a, b))

	//返回第一次出现字符串索引
	fmt.Println(strings.Index(con, "s"))
}
