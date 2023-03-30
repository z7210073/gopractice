package main

import (
	"20230110/package/model"
	"fmt"
)

// 导入自定义的包需要到项目根目录输入 go mod init <项目名> 命令生成 go.mod文件

func main() {
	fmt.Println(model.HeroName)

}
