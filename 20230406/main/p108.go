package main

import "fmt"

func main() {
	//goto 跳转到指定的行 ,不主张使用
label:
	fmt.Println(1)

	for i := 0; i < 10; i++ {
		if i == 1 {
			goto label //无限输出1
		}
	}

}
