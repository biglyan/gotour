package main

import "fmt"

type data struct {
	val int
}

func (pData *data) set(num int) {
	pData.val = num
}

func (pData *data) show() {
	fmt.Println(pData.val)
}

func main() {
	pData := &data{4}
	pData.set(5)
	pData.show()
}