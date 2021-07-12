package utils

import (
	"fmt"
	"os"
)

func Exit() {
	os.Exit(1)
}

func Print(text string) {
	fmt.Print(text)
}

func Scanln(text interface{}) {
	fmt.Scanln(text)
}

//InpAndOut 输入 + 输出函数
func InpAndOut(t string, c interface{}) {
	Print(t)
	Scanln(c)
}
