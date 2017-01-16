package main

import (
	"fmt"
	"time"
)

func show() {
	for {
		fmt.Print("child ")
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	go show()
	for {
		fmt.Print("parent ")
		time.Sleep(2000 * time.Millisecond)
	}
}