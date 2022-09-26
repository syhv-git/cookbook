package main

import (
	"cookbook/file/forensics"
	"fmt"
)

func main() {
	fileForensics()
}

func fileForensics() {
	// enumerate paths and return slice of all accessible files sorted based on params
	res := forensics.Enumerate("mod", true, "/home/scott")
	for x := 0; x < 75; x++ {
		fmt.Println(res[x])
	}

	// extract contents at paths and write it to dest
	forensics.ExtractCopy("output.txt", "")
}
