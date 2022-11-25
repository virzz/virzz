package image

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"math"
	"os"
	"path"
	"strings"

	"github.com/mozhu1024/virzz/logger"
	"github.com/mozhu1024/virzz/tools/downloader"
	"github.com/skip2/go-qrcode"
	decode "github.com/tuotoo/qrcode"
)

// type qrcodeRecoveryLevel qrcode.RecoveryLevel
// type qrcodeRecoveryLevels struct {
// 	Low     qrcodeRecoveryLevel
// 	Medium  qrcodeRecoveryLevel
// 	High    qrcodeRecoveryLevel
// 	Highest qrcodeRecoveryLevel
// }

func printQrcodeToTerminal(data [][]bool) string {
	var builder strings.Builder
	builder.Grow((15*len(data) + 1) * len(data))
	for _, row := range data {
		for _, col := range row {
			if col {
				builder.WriteString("\033[48;5;0m  \033[0m")
			} else {
				builder.WriteString("\033[48;5;7m  \033[0m")
			}
		}
		builder.WriteByte('\n')
	}
	return builder.String()
}

// ZeroOneToQrcode -
func ZeroOneToQrcode(s string, re bool, dst string) (string, error) {
	var flag byte = '1'
	if re {
		flag = '0'
	}
	length := int(math.Sqrt(float64(len(s))))
	if length < 1 {
		return "", fmt.Errorf("input data error")
	}
	logger.DebugF("Image Flag = %c Image Size = %d", flag, length)
	if dst == "-" {
		// var  [][]bool
		data := make([][]bool, length)
		for i := range data {
			data[i] = make([]bool, length)
		}
		for y := 0; y < length; y++ {
			for x := 0; x < length; x++ {
				data[x][y] = s[y*length+x] == flag
			}
		}
		return printQrcodeToTerminal(data), nil
	}
	img := image.NewRGBA(image.Rect(0, 0, length, length))
	draw.Draw(img, img.Bounds(), &image.Uniform{image.White}, img.Bounds().Min, draw.Src)
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

// ReadQrcode -
func ReadQrcode(src string, terminal ...bool) (string, error) {
	var (
		file     *os.File
		filename string
	)
	if strings.HasPrefix(src, "http") {
		filename = path.Join(os.TempDir(), "qrcode.png")
		if err := downloader.SigleFetch(src, filename); err != nil {
			return "", err
		}
	} else {
		filename = src
	}
	logger.Debug("filename: ", filename)
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	qrMatrix, err := decode.Decode(file)
	if err != nil {
		logger.Debug(err)
		return "", err
	}
	if len(terminal) > 0 && terminal[0] {
		return printQrcodeToTerminal(qrMatrix.Points), nil
	}
	return qrMatrix.Content, nil
}

func WriteQrcode(src, dst string) (string, error) {
	var qr *qrcode.QRCode
	var err error
	qr, err = qrcode.New(src, qrcode.RecoveryLevel(qrcode.Medium))
	if err != nil {
		return "", err
	}
	if qr == nil {
		return "", fmt.Errorf("src error")
	}
	if dst == "-" {
		data := qr.Bitmap()
		return printQrcodeToTerminal(data), nil
	}
	err = qr.WriteFile(256, dst)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Generate [%s] success", dst), nil
}
