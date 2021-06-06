package main

import "fmt"

type module string

var (
	A module = "a"
	B module = "b"
	C module = "c"
	D module = "d"
)

func main ()  {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
}
