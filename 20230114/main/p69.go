package main

import "fmt"

func main() {
	//获取终端输入
	//从控制台获取 年龄 薪水  姓名 是否通过考试
	var name string
	var age int
	var sal float64
	var isPass bool
	//// 例子1
	//fmt.Println("请输入姓名")
	//fmt.Scanln(&name)
	//fmt.Println("请输入年龄")
	//fmt.Scanln(&age)
	//fmt.Println("请输入薪水")
	//fmt.Scanln(&sal)
	//fmt.Println("请输入是否通过")
	//fmt.Scanln(&isPass)

	// 例子2
	fmt.Println("请输入年龄 薪水  姓名 是否通过考试使用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(sal)
	fmt.Println(isPass)

}
