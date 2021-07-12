package main

import "fmt"

func main() {
	printStar()
}

var v = `
     *
    ***
   *****
  *******
 *********
***********
***********
 *********
  *******
   *****
    ***
     *`

// printStar 打印**
func printStar() {
	// 复杂度o(1), 最优解
	fmt.Println(v)
}
