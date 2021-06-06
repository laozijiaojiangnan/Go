package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// 类型转换
func main() {
	s := "123"
	s1 , _ := strconv.Atoi(s)
	fmt.Println(s1,reflect.TypeOf(s1))
}
