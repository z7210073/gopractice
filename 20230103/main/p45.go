package main

import "fmt"

func main() {
	//go 中字符串不可变
	var str = "hello"

	//str[0] = '1'   cannot assign to str[0] (value of type byte)

	fmt.Println(str)
	// 除了转义字符 \  反引号    ` `  也可以进行输出

	fmt.Println("\\")
	fmt.Println(`\`)

	//字符串拼接
	var str2 = "hello" + "world"

	//换行加号在上面
	var str3 = "hello" +
		"world" + "sss"
	str2 += "sss"

	fmt.Println(str2)
	fmt.Println(str3)

}
