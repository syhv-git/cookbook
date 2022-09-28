package utility

import (
	cmd "cookbook"
	"image"
	"image/jpeg"
	"io"
	"math/rand"
)

func GenerateRandomImage(v bool, x, y int, dst io.Writer) {
	if v {
		cmd.Log(v, "- Generating random noise as image")
	}

	i := image.NewRGBA(image.Rect(0, 0, x, y))
	for c := 0; c < x*y; c++ {
		o := 4 * c
		i.Pix[0+o] = uint8(rand.Intn(256))
		i.Pix[1+o] = uint8(rand.Intn(256))
		i.Pix[2+o] = uint8(rand.Intn(256))
		i.Pix[3+o] = 255
	}

	cmd.Log(v, "- Encoding the generated RGB to JPEG")
	err := jpeg.Encode(dst, i, nil)
	if err != nil {
		cmd.Fatal("## " + err.Error())
	}
}
