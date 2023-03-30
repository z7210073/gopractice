package main

import (
	"fmt"
	"strconv"
)

func main() {

	//int to string

	s := strconv.Itoa(1)
	fmt.Printf("%T \n", s)

	// string 转其他类型
	var str string = "false"
	var b bool
	b, _ = strconv.ParseBool(str) //多值返回，下划线表示忽略第二个返回值

	fmt.Printf("%T \n ", b)

	var strToInt string = "123"
	var i int64
	i, _ = strconv.ParseInt(strToInt, 10, 64)

	fmt.Printf("%T \n", i)

	var strToFloat string = "123.23"
	var v float64

	v, _ = strconv.ParseFloat(strToFloat, 64)
	fmt.Printf("%T,%v", v, v)

}
