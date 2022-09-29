package tests

import (
	"bytes"
	cmd "cookbook"
	"cookbook/file/forensics"
	"os"
	"testing"
)

func TestEnumeration(t *testing.T) {
	ret := forensics.Enumerate(true, "size", true, "/home/scott/Desktop")
	if len(ret) < 1 {
		t.Fatal("## Failed to enumerate project root")
	}
	for x := 0; x < 50; x++ {
		cmd.Log(true, "%d", ret[x].Size())
	}
}

func TestExtractor(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	forensics.Extractor(true, buf, "types_test.go")

	a, _ := os.ReadFile("types_test.go")
	if string(buf.Bytes()) != string(a) {
		t.Error("## Error when extracting file contents")
	}
}

func TestExtractCopy(t *testing.T) {
	forensics.ExtractCopy(true, "tmp/extract.txt", "types_test.go")
	defer os.RemoveAll("tmp")

	f, err := os.ReadFile("tmp/extract.txt")
	if err != nil {
		t.Fatal("## " + err.Error())
	}
	a, _ := os.ReadFile("types_test.go")
	if string(f) != string(a) {
		t.Error("## Error when extracting file contents")
	}
}
