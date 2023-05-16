package main

import (
	"fmt"
)

func sums(n1 int) int {
	n1++                  //11
	defer fmt.Println(n1) //11
	n1++                  //12
	return n1             //12
}
func main() {
	//defer 将将值拷贝到栈中
	fmt.Println(sums(10))

}
