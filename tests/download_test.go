package tests

import (
	"github.com/syhv-git/cookbook/network/utility"
	"os"
	"testing"
)

func TestHTTPDownloader(t *testing.T) {
	url := "https://en.wikipedia.org/wiki/File:Image_created_with_a_mobile_phone.png"
	dst := "test_download.png"

	utility.Download(dst, url)
	info, err := os.Stat(dst)
	if err != nil {
		t.Fatal("## " + err.Error())
	}
	if info.Size() < 1 {
		t.Error("## Error when download file")
	}

	os.Remove(dst)
}
