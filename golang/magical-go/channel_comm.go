package main

import "fmt"

func show(c chan int) {
	for {
		data := <-c
		if data == 1 {
			fmt.Print("receive ")
		}
	}
}

func main() {
	c := make(chan int)
	go show(c)
	for {
		c <- 1
		fmt.Print("send ")
	}
}