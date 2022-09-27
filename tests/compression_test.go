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
	os.Remove(dst)
}
