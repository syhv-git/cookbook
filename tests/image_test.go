package tests

import (
	"github.com/syhv-git/cookbook/file/utility"
	"os"
	"testing"
)

func TestGenerateRandomIMG(t *testing.T) {
	dst := "tmp.jpg"
	f, err := os.Create(dst)
	if err != nil && !os.IsExist(err) {
		t.Fatal("## " + err.Error())
	}
	utility.GenerateRandomImage(true, 400, 400, f)
	info, _ := f.Stat()
	if info.Size() < 1 {
		t.Error("## Error when generating image")
	}
	os.Remove(dst)
}
