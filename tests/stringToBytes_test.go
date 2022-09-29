package tests

import (
	"bytes"
	"github.com/syhv-git/cookbook/file/utility"
	"testing"
)

func TestStringToBytes(t *testing.T) {
	b := utility.StringToBytes(true, 32, "Hello World!\n")
	if !bytes.Equal(b, []byte{'l', 'l', 'e', 'H', 'o', 'W', ' ', 'o', '!', 'd', 'l', 'r', '\x00', '\x00', '\x00', '\n'}) {
		t.Error("## Error when converting string to bytes in little-endian with 32-bit registers")
	}
}
