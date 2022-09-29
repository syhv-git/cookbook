package tests

import (
	"cookbook/file/forensics"
	"cookbook/file/utility"
	"os"
	"testing"
)

func TestCreateSteganographicFromArchive(t *testing.T) {
	dst := "stego_test.jpg"
	src := "test_stego.zip"
	jpg := "image_test.jpg"
	utility.CompressNew(true, src, "../.gitignore", "../file/types.go", "types_test.go")

	forensics.CreateSteganographicFromArchive(true, dst, src, jpg)
	info, err := os.Stat(dst)
	if err != nil {
		t.Fatal("## " + err.Error())
	}
	if info.Size() < 1 {
		t.Error("## Error when creating steganographic archive file")
	}

	// TestDetectArchiveFromImage
	b := forensics.DetectArchiveFromImage(true, dst)
	if !b {
		t.Error("## Error when detecting archive in image file")
	}
	
	// Test decompress archive (for Zip archives, GZip doesnt seem to work)
	//utility.Decompress(true, "test", dst)

	os.Remove(dst)
	os.Remove(src)
	//os.Remove(jpg)
}
