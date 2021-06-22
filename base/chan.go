package main

import (
	"fmt"
)

func main() {
	// 用来收数据的通道
	var chans [10] chan<- int
	var chans2 [10] <-chan bool
	for i:=0; i < 10; i++ {
		chans[i], chans2[i] = CreateWorker(i)
	}

	for idx, c := range chans {
		c <- idx
	}

	for _, c := range chans2 {
		<-c
	}
}

func worker(id int,c chan int, done chan bool) {
	for {
		n := <-c
		fmt.Println("worker id:", id, n)
		done <- true
	}
}

func CreateWorker(id int) (chan<- int, <-chan bool){
	c := make(chan int)
	d := make(chan bool)
	go worker(id, c, d)
	return c, d
}
