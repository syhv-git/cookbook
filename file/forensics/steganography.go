package forensics

import (
	cmd "cookbook"
	"io"
	"os"
)

// CreateSteganographicFromArchive creates an image file that behave both as an image and archive.
// dst is the path for the new image archive and src is the path to the archive.
// jpg is an optional path to an image file
func CreateSteganographicFromArchive(v bool, dst, src, jpg string) {
	if src == "" {
		cmd.Fatal("## Source file was not provided")
	}
	if dst == "" {
		cmd.Fatal("## Destination file was not provided")
	}
	if jpg == "" {
		cmd.Fatal("## Source image file was not provided")
	}
	j, err := os.Open(jpg)
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}
	defer j.Close()

	a, err := os.Open(src)
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}
	defer a.Close()

	s, err := os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}
	defer s.Close()

	if _, err = io.Copy(s, j); err != nil {
		cmd.Fatal("## " + err.Error())
	}
	if _, err = io.Copy(s, a); err != nil {
		cmd.Fatal("## " + err.Error())
	}
}

//func DetectArchiveFromImage(v bool, src string) {
//
//}
