package main

import (
	"fmt"
	"flag"
)

var _version_ = ""
var _branch_ = ""
var _commitId_ = ""
var _buildTime_ = ""

func main() {
	var version bool
	flag.BoolVar(&version, "v", false, "-v")
	flag.Parse()

	if version {
		fmt.Printf("Version: %s, Branch: %s, Build: %s, Build time: %s\n",
			_version_, _branch_, _commitId_, _buildTime_)
	}
}