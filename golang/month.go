package main

import (
	"fmt"
	"time"
)

func main() {
	year, month, day := time.Now().Date()
	fmt.Printf("year %v, month %v, day %v", year, month, day)
}