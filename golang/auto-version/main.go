package main

import (
	"flag"
	"fmt"
)

var (
	GitTag = "2017.01.24.1343.release"
	BuildTime = "2017-01-24T13:44:50+0800"
)

func main() {
	version := flag.Bool("v", false, "version")
	flag.Parse()

	if *version {
		fmt.Println("Git Tag: " + GitTag)
		fmt.Println("Build Time: " + BuildTime)
	}
}