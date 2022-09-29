package main

import (
	"cookbook/cmd"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		cmd.StartUp()
	}
}
