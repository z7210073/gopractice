package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//为随机数设置一个rand种子
	rand.Seed(time.Now().Unix())

	for {
		n := rand.Intn(100) + 1
		if n == 99 {
			break
		}
		fmt.Println(n)
	}
}
