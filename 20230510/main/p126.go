package main

import (
	"fmt"
	"strings"
)

// 匿名函数和sufix构成闭包
func makeSufixx(sufix string) func(string) string {

	return func(file string) string {
		if strings.HasSuffix(file, sufix) {
			return file
		} else {
			return file + sufix

		}
	}
}
func main() {
	f := makeSufixx(".jpg")
	fmt.Println(f("1.jpg"))
	fmt.Println(f("2"))
}
