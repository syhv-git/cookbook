package main

import (
	"cookbook/file/forensics"
	"fmt"
)

func main() {
	fileForensics()
}

func fileForensics() {
	res := forensics.Enumerate("size", true, "/home/scott/Desktop")
	for _, x := range res {
		fmt.Println(x)
	}
}
