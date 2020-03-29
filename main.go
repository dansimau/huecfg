package main

import (
	"os"

	"github.com/dansimau/huecfg/cmd"
)

func main() {
	os.Exit(cmd.Run(os.Args[1:]))
}
