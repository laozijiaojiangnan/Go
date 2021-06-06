package main

import "fmt"

var (
	v1 = "123"
)

func main()  {
	hello()
	printLine()
	中文_()
}

// 测试中文
func 中文_()  {
	a, b, c := 1, "1", true
	fmt.Println(a, b, c)
}

// hello 开始程序
func hello()  {
	fmt.Println("hello\"go\"")
}

// _line 生成分割线
func printLine()  {
	fmt.Println("--------------------------------------------")
}

