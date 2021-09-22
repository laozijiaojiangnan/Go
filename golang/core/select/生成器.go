package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s1, s2, s3 := generateTask("s1"), generateTask("s2"), generateTask("s3")
	c := _print(s1, s2, s3)
	for {
		fmt.Println(<-c)
	}
}

func _task(name string, c chan string) {
	i := 0
	for {
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		c <- fmt.Sprintf("task name:%v, content:%v", name, i)
		i++
	}
}

func generateTask(name string) chan string {
	c := make(chan string)
	go _task(name, c)
	return c
}

func _print(chans ...chan string) chan string{
	nc := make(chan string)
	for _, c := range chans {
		go func(c chan string) {
			for {
				nc <- <-c
			}
		}(c)
	}
	return nc
}
