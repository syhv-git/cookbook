package main

import (
	"cookbook/cmd"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		cmd.StartUp()
		return
	}

	switch os.Args[1] {
	case "help":
		cmd.Help("basic")
	default:
		cmd.Fatal("## Unknown argument: %s\n\tRun 'cookbook help' for documentation")
	}
}
