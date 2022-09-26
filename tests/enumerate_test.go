package tests

import (
	"cookbook/file/forensics"
	"log"
	"os"
	"testing"
)

func TestEnumeration(t *testing.T) {
	ret := forensics.Enumerate("size", true, "/home/scott/Desktop")
	if len(ret) < 1 {
		t.Fatal("Failed to enumerate project root")
	}
	for x := 0; x < 50; x++ {
		log.Println(ret[x].Size())
	}
}

func TestExtraction(t *testing.T) {
	forensics.ExtractCopy("tmp/extract.txt", "types_test.go")
	defer os.RemoveAll("tmp")

	f, err := os.ReadFile("tmp/extract.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
	a, _ := os.ReadFile("types_test.go")
	if string(f) != string(a) {
		t.Error("Error when extracting file contents")
	}
}
