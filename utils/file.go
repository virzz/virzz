package utils

import (
	"bytes"
	"io"
	"os"
	"path"

	"github.com/virzz/logger"
)

// Copy file from [src] to [dest] by little buffer
func CopyFile(dest, src string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	os.MkdirAll(path.Dir(dest), 0755)

	out, err := os.OpenFile(dest, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer out.Close()
	inStat, _ := in.Stat()
	outStat, _ := out.Stat()

	var inBuf, outBuf = make([]byte, 1024), make([]byte, 1024)
	in.Read(inBuf)
	out.Read(outBuf)
	logger.DebugF(
		"src: %s %d %s\ndst: %s %d %s\n",
		inStat.Name(), inStat.Size(), string(inBuf),
		outStat.Name(), outStat.Size(), string(outBuf),
	)
	if inStat.Name() == outStat.Name() &&
		inStat.Size() == outStat.Size() &&
		bytes.Equal(inBuf, outBuf) {
		logger.Warn("There are same file, stop copy")
		return nil
	}

	in.Seek(0, 0)
	out.Seek(0, 0)
	data := make([]byte, 1024*1024) // 1M
	for {
		n, e := in.Read(data)
		if e == io.EOF || n == 0 {
			return nil
		}
		out.Write(data[:n])
	}
}
