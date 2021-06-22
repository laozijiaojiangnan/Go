package main

import "fmt"

func main() {
	_test()
}

func _test() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7}
	b := a[0:4]
	b = append(b, 4444)
	fmt.Println(a)
	fmt.Println(b)
}
