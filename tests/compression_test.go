package tests

import (
	"github.com/syhv-git/cookbook/cmd"
	"github.com/syhv-git/cookbook/file/utility"
	"os"
	"testing"
)

func TestNewCompressionAndDecompression(t *testing.T) {
	dst := "test.tar.gz"
	utility.CompressNew(true, dst, "../.gitignore", "../file/types.go", "types_test.go")
	f, err := os.Stat(dst)
	if err != nil {
		t.Fatal("## " + err.Error())
	}
	if f.Name() != dst {
		t.Error("## Error when creating compressed file")
	}

	// * Testing decompress function
	cmd.Log(true, "* Testing decompression")
	utility.Decompress(true, "", dst)
	d, err := os.ReadDir("test")
	if err != nil {
		t.Fatal("## " + err.Error())
	}
	if len(d) < 1 {
		t.Error("## Error when decompressing file")
	}

	os.Remove(dst)
	os.RemoveAll("test")
}

func TestDecompression(t *testing.T) {
	cmd.Log(true, "* Testing decompression")
	utility.Decompress(true, "", "stego_test.jpg")
	d, err := os.ReadDir("test")
	if err != nil {
		t.Fatal("## " + err.Error())
	}
	if len(d) < 1 {
		t.Error("## Error when decompressing file")
	}
}
