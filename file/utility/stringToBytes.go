package utility

import (
	"github.com/syhv-git/cookbook/cmd"
	"os"
)

// StringToBytes converts a string to raw byte form.
// The variable l defines register size for little-endian. 0 value indicates big-endian
func StringToBytes(v bool, l int, s string) (bs []byte) {
	if s == "" {
		cmd.Fatal("## No string provided")
	}

	for _, b := range s {
		bs = append(bs, byte(b))
	}
	cmd.Log(v, "- Printing the string as bytes")
	switch l {
	case 16:
		bs = littleEndian(bs, 2)
	case 32:
		bs = littleEndian(bs, 4)
	case 64:
		bs = littleEndian(bs, 8)
	}

	if v {
		if _, err := os.Stdout.Write(bs); err != nil {
			cmd.Fatal("## " + err.Error())
		}
		os.Stdout.Write([]byte{'\n'})
	}
	return
}

func littleEndian(bs []byte, r int) []byte {
	dif := len(bs) % r
	size := len(bs) + dif
	res := make([]byte, size)

	for i := 0; i < size/r; i += r {
		for j, k := i*r, i*r+r; j < k; j, k = j+1, k-1 {
			if k < len(bs) {
				res[j], res[k] = bs[k], bs[j]
			} else {
				res[j], res[k] = '\x00', bs[j]
			}
		}
	}
	return res
}
