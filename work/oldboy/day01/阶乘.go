package main

import "fmt"

func main() {
	fmt.Println(f1())
}

// f1 阶乘
func f1() int {
	sum := 0
	for i := 1; i < 11; i++ {
		tmp := 1
		for k := 1; k < i+1; k++ {
			tmp *= k
		}
		sum += tmp
	}
	return sum
}
