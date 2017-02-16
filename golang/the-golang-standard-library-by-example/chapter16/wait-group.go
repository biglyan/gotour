package main

import (
	"fmt"
	"sync"
)

func main() {
	// 一个goroutine需要等待一批goroutine执行完毕以后才继续执行，
	// 那么这种多线程等待的问题就可以使用WaitGroup了。
	wp := new(sync.WaitGroup)
	wp.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("done ", i)
			wp.Done()
		}()
	}

	wp.Wait()
	fmt.Println("wait end")
}
