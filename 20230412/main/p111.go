package main

import (
	"20230412/package/db"
	"20230412/package/utils"
	"fmt"
)

func main() {
	//为完成某一功能的集合叫函数
	//go 分自定义函数和系统函数
	//函数 并非方法
	//包可以控制变量和函数访问范围
	fmt.Println(utils.Calc(1, 1))
	fmt.Println(db.Db("123"))

}
