package main

import (
	"fmt"
)

func main() {
	fmt.Println(FibOV2(7)) // 13
}

// FibO 清晰版
func FibO(count int) int {
	j, k, tmp := 0, 1, 0
	for i := 0; i < count; i++ {
		tmp = j + k
		j = k
		k = tmp
	}
	return j
}

// FibOV2 优化版
func FibOV2(count int) int {
	k, j := 0, 1
	for i := 0; i < count; i++ {
		k, j = j, j+k
	}
	return k
}
