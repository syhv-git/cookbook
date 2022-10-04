package forensics

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"github.com/syhv-git/cookbook/cmd"
	"os"
)

func Checksum(file string) {
	b, err := os.ReadFile(file)
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}
	cmd.Log(true, "- Logging hashes for: %s", file)
	cmd.Log(true, "* MD5: %x", md5.Sum(b))
	cmd.Log(true, "* SHA1: %x", sha1.Sum(b))
	cmd.Log(true, "* MD5: %x", sha256.Sum256(b))
	cmd.Log(true, "* MD5: %x", sha512.Sum512(b))
}

func VerifyChecksum(file string, sum string) bool {
	b, err := os.ReadFile(file)
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}

	m := md5.Sum(b)
	if bytes.Equal(m[:], []byte(sum)) {
		return true
	}

	s1 := sha1.Sum(b)
	if bytes.Equal(s1[:], []byte(sum)) {
		return true
	}

	s256 := sha256.Sum256(b)
	if bytes.Equal(s256[:], []byte(sum)) {
		return true
	}

	s512 := sha512.Sum512(b)
	if bytes.Equal(s512[:], []byte(sum)) {
		return true
	}
	return false
}
