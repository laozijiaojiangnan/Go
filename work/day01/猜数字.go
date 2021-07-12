package main

import (
	"fmt"
)

func main() {
	guessNumber()
}

var (
	number = 66
	inp    int
	text   = `
猜数字游戏
规则如下
1. 输入非数字统一为0
2. 只有三次机会，老弟`
)

// guessNumber 猜数字
func guessNumber() {
	fmt.Println(text)
	c := 3
	for {
		if c <= 0 {
			fmt.Println("次数用光啦!")
			break
		}

		// 接受输入
		fmt.Printf("还有%v次机会，请输入：", c)
		fmt.Scanln(&inp)

		if inp > number {
			fmt.Println("太大啦!")
		} else if inp < number {
			fmt.Println("太小啦!")
		} else {
			fmt.Println("\\^o^/答对啦")
			break
		}
		c -= 1
	}

}
