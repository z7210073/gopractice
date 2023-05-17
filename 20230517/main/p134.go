package main

import (
	"fmt"
	"strings"
)

func main() {
	//返回最后一次索引
	var con string = "ssIamsss d"
	fmt.Println(strings.LastIndex(con, "s"))
	//字符串替换 n表示替换几个， n=-1 表示全部替换
	con2 := strings.Replace(con, "s", "i", 2)
	fmt.Println(con2)
	//字符串拆分
	var str string = "hello,world,ok"
	fmt.Println(strings.Split(str, ","))
	//大小写转换
	fmt.Println(strings.ToLower("WWWasddsaA"))
	fmt.Println(strings.ToUpper("ssdsaADSA"))

}
