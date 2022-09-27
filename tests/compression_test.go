package tests

import (
	"cookbook/file/utility"
	"os"
	"testing"
)

func TestNewCompression(t *testing.T) {
	utility.CompressNew(true, "test.tar.gz", "../.gitignore", "../file/types.go", "types_test.go")
	f, err := os.Stat("test.tar.gz")
	if err != nil {
		t.Fatal(err.Error())
	}
	if f.Name() != "test.tar.gz" {
		t.Error("Error when creating compressed file")
	}
}
