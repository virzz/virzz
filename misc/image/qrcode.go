package image

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"math"
	"os"

	cm "github.com/virink/virzz/common"
)

// ZeroOneToQrcode -
func ZeroOneToQrcode(s string, re bool, dst string) (string, error) {
	var flag byte = '1'
	if re {
		flag = '0'
	}
	cm.Logger.Debug("Image Flag = %c", flag)
	length := int(math.Sqrt(float64(len(s))))
	if length < 1 {
		return "", fmt.Errorf("input data error")
	}
	cm.Logger.Debug("Image Size = %d", length)
	img := image.NewRGBA(image.Rect(0, 0, length, length))
	draw.Draw(img, img.Bounds(), &image.Uniform{image.White}, image.ZP, draw.Src)
	// draw.Draw(img)
	for y := 0; y < length; y++ {
		for x := 0; x < length; x++ {
			if s[y*length+x] == flag {
				img.Set(x, y, image.Black)
			}
		}
	}
	// TODO: Scale (graphics.Scale())
	f, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer f.Close()
	err = png.Encode(f, img)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Generate [%s] success", dst), nil
}
