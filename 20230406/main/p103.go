package main

import "fmt"

func main() {
	// for可以设置标签 可以通过break +标签停止对应的循环
label2:
	for {

		for {
			for i := 0; i < 10; i++ {
				fmt.Println(i)
				continue label2
				if i == 4 {
					fmt.Println("跳出所有循环")
					break label2

				}
			}
		}
	}
}
