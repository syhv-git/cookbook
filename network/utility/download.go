package utility

import (
	"io"
	"net/http"
	"os"
)

func Download(dst, url string) {
	f, err := os.Create(dst)
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}
	defer f.Close()

	r, err := http.Get(url)
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}
	defer r.Body.Close()

	if _, err = io.Copy(f, r.Body); err != nil {
		cmd.Fatal("## " + err.Error())
	}
}
