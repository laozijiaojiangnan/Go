package main

import (
	"fmt"
	"time"
)

func generate() chan int{
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Second)
			out <- i
			i++
		}
	}()
	return out
}

func main()  {
	//c1, c2 := generate(), generate()
	//for {
	//	select {
	//	case n := <-c1:
	//		fmt.Println("received from c1:", n)
	//	case n := <-c2:
	//		fmt.Println("received from c2:", n)
	//	default:
	//		fmt.Println("no value received")
	//		time.Sleep(time.Second)
	//	}
	//}
	num := 0
	go func() {
		for i:= 0; i < 1000; i++{
			num++
		}
	}()
	go func() {
		for i:= 0; i < 1000; i++{
			num--
		}
	}()

	time.Sleep(time.Second)
	fmt.Println(num)
}
