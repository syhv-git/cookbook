package tests

import (
	"cookbook/file/utility"
	"os"
	"testing"
)

func TestNewCompression(t *testing.T) {
	dst := "test.tar.gz"
	utility.CompressNew(true, dst, "../.gitignore", "../file/types.go", "types_test.go")
	f, err := os.Stat(dst)
	if err != nil {
		t.Fatal(err.Error())
	}
	if f.Name() != dst {
		t.Error("Error when creating compressed file")
	}

	utility.Decompress(true, "", dst)
	d, err := os.ReadDir("test")
	if err != nil {
		t.Fatal(err.Error())
	}
	if len(d) < 1 {
		t.Error("Error when decompressing file")
	}

	os.Remove(dst)
	os.RemoveAll("test")
}
