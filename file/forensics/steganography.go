package forensics

import (
	"bufio"
	"bytes"
	"cookbook/cmd"
	"io"
	"os"
)

// CreateSteganographicFromArchive creates an image file that behave both as an image and archive.
// dst is the path for the new image archive and src is the path to the archive.
// jpg is an optional path to an image file
func CreateSteganographicFromArchive(v bool, dst, src, jpg string) {
	cmd.Log(v, "*** Starting Steganography")
	defer cmd.Log(v, "*** Ending Steganography")
	if src == "" {
		cmd.Fatal("## No Source file provided")
	}
	if dst == "" {
		cmd.Fatal("## No Destination file provided")
	}
	if jpg == "" {
		cmd.Fatal("## No Source image file provided")
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

	cmd.Log(v, "- Creating destination file: %s", dst)
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

var signatures = [][]byte{
	{'\x37', '\x7A', '\xBC', '\xAF', '\x27', '\x1C'}, // 7Zip
	{'\x42', '\x5a', '\x68'},                         // BZip2
	{'\x7F', '\x45', '\x4C', '\x46'},                 // ELF
	{'\x4d', '\x5a'},                                 // EXE
	{'\x1f', '\x8b'},                                 // GZip
	{'\x1f', '\xa0'},                                 // LZH
	{'\x4C', '\x5A', '\x49', '\x50'},                 // LZip
	{'\x1f', '\x9d'},                                 // LZW
	{'\x52', '\x61', '\x72', '\x21', '\x1A', '\x07'}, // RAR
	{'\x75', '\x73', '\x74', '\x61', '\x72'},         // Tar
	{'\xEF', '\xBB', '\xBF'},                         // TXT
	{'\x50', '\x4b', '\x03', '\x04'},                 // Zip
}

func DetectArchiveFromImage(v bool, src string) bool {
	cmd.Log(v, "*** Starting Steganography detection")
	defer cmd.Log(v, "*** Ending Steganography detection")

	f, err := os.Open(src)
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}
	r := bufio.NewReader(f)

	info, _ := f.Stat()
	for i := int64(0); i < info.Size(); i++ {
		b, err := r.ReadByte()
		if err != nil {
			cmd.Fatal("## " + err.Error())
		}

		for _, s := range signatures {
			if b == s[0] {
				c := make([]byte, len(s)-1)
				c, err := r.Peek(len(s) - 1)
				if err != nil {
					cmd.Fatal("## " + err.Error())
				}
				if bytes.Equal(c, s[1:]) {
					cmd.Log(v, "* Detected embedded file signature (little endian): 0x%x", s[:])
					return true
				}
			}
		}
	}
	return false
}
