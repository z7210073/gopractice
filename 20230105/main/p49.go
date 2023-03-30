package main

import (
	"fmt"
	"strconv"
)

func main() {

	//string 与其他类型互相转换

	//转string 方法1
	var num1 int64 = 99
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'
	var newStr1 string
	newStr1 = fmt.Sprintf("%d", num1) //第一个为本身的数据类型第二个为要转的参数
	fmt.Printf(newStr1+" type is %T\n", newStr1)
	newStr1 = fmt.Sprintf("%f", num2) //第一个为本身的数据类型第二个为要转的参数
	fmt.Printf(newStr1+" type is %T\n", newStr1)
	newStr1 = fmt.Sprintf("%t", b) //第一个为本身的数据类型第二个为要转的参数
	fmt.Printf(newStr1+" type is %T\n", newStr1)
	newStr1 = fmt.Sprintf("%c", myChar) //第一个为本身的数据类型第二个为要转的参数
	fmt.Printf(newStr1+" type is %T\n", newStr1)

	//转string 方法1
	var newStr2 string
	newStr2 = strconv.FormatInt(num1, 10)
	fmt.Printf(newStr2+" type is %T\n", newStr2)
	newStr2 = strconv.FormatFloat(num2, 'f', 10, 64) //第一个为要转的值，‘f’要代表转成的格式，10代表精度，64代表小数为float64
	fmt.Printf(newStr2+" type is %T\n", newStr2)
	newStr2 = strconv.FormatBool(b)
	fmt.Printf(newStr2+" type is %T\n", newStr2)
	//newStr2 = strconv.FormatBool(b)
	//fmt.Printf(newStr2+" type is %T\n", newStr2)
}
