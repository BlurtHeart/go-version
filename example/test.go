package main

import (
	"flag"

	"github.com/blurty/go-version"
)

func main() {
	showVersion := flag.Bool("v", false, "version info")
	flag.Parse()

	if *showVersion == true {
		version.PrintVersion()
	}
}
