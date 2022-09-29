package tests

import (
	"cookbook/file/forensics"
	"cookbook/file/utility"
	"os"
	"testing"
)

func TestCreateSteganographicFromArchive(t *testing.T) {
	dst := "stego_test.jpg"
	src := "test_stego.tar.gz"
	utility.CompressNew(true, src, "../.gitignore", "../file/types.go", "types_test.go")

	//f, err := os.OpenFile("image_test.jpg", os.O_CREATE|os.O_TRUNC, 0666)
	//utility.GenerateRandomImage(true, 400, 400, f)
	//f.Close()

	forensics.CreateSteganographicFromArchive(true, dst, src, "image_test.jpg")
	info, err := os.Stat(dst)
	if err != nil {
		t.Fatal("## " + err.Error())
	}
	if info.Size() < 1 {
		t.Error("## Error when creating steganographic archive file")
	}

	os.Remove(dst)
	os.Remove(src)
	os.Remove("image_test.jpg")
}
