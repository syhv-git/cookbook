package tests

import (
	"cookbook/file/utility"
	"os"
	"testing"
)

func TestNewCompression(t *testing.T) {
	utility.CompressNew(true, "test.zip", "../.gitignore", "../file/types.go", "types_test.go")
	f, err := os.Stat("test.zip")
	if err != nil {
		t.Fatal(err.Error())
	}
	if f.Name() != "test.zip" {
		t.Error("Error when creating compressed file")
	}
}
