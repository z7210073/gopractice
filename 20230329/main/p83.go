package main

import "fmt"

func main() {
	a := 1
	if a > 1 {
		fmt.Println(1)
	} else if a == 1 {
		fmt.Println(1)
	} else { //else 部分可以没有
		fmt.Println(2)
	}
}
