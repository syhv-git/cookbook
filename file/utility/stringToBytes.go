package utility

import (
	"cookbook/cmd"
	"os"
)

func StringToBytes(v, l bool, s string) (bs []byte) {
	if s == "" {
		cmd.Fatal("## No string provided")
	}
	for _, b := range s {
		bs = append(bs, byte(b))
	}
	if v {
		cmd.Log(v, "- Printing the string as bytes")
		if !l {
			if _, err := os.Stdout.Write(bs); err != nil {
				cmd.Fatal("## " + err.Error())
			}
		} else {
			o := len(bs) % 4
			if o != 0 {

			}
			// every 4 bytes are reversed
			for i := 0; i < len(bs)-1; i += 4 {
				bs[i], bs[i+1], bs[i+2], bs[i+3] = bs[i+3], bs[i+2], bs[i+1], bs[i]
			}
			if _, err := os.Stdout.Write(bs); err != nil {
				cmd.Fatal("## " + err.Error())
			}
		}
	}
	return
}
