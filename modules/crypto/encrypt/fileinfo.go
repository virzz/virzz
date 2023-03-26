package encrypt

import (
	"bytes"
	"encoding/binary"
	"io"
	"os"

	"github.com/pkg/errors"
	"github.com/virzz/logger"
)

var ErrNotEncryptedByVirzz = errors.New("Not Encrypted By virzz")

type EncryptFileInfo struct {
	Size       uint32
	Block      uint32
	Ext        []byte
	IsCompress bool
}

func (e *EncryptFileInfo) Data() (int, []byte) {
	/*
		Len	Data
		n	Ext (n<=32)
		8	Block
		8	Size
		1	IsCompress
		1	uint8(len(e.ext))
		5	VIRZZ
		pngaaaabbbbcdVIRZZ
	*/
	if len(e.Ext) > 32 {
		// 不支持长度超过32字节
		e.Ext = e.Ext[:32]
	}
	var buf bytes.Buffer
	buf.Write(e.Ext)
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(e.Block))
	buf.Write(bs)
	binary.LittleEndian.PutUint32(bs, uint32(e.Size))
	buf.Write(bs)
	if e.IsCompress {
		buf.WriteByte(1)
	} else {
		buf.WriteByte(0)
	}
	buf.WriteByte(uint8(len(e.Ext)))
	buf.WriteString("VIRZZ")
	return buf.Len(), buf.Bytes()
}

func (e *EncryptFileInfo) Read(data []byte) (*EncryptFileInfo, error) {
	if !bytes.HasSuffix(data, []byte("VIRZZ")) {
		return nil, ErrNotEncryptedByVirzz
	}
	l := len(data)
	extSize := int(data[l-6])
	// if len(data) < extSize {
	// 	return nil
	// }
	e.IsCompress = data[l-7] == 1
	e.Ext = data[l-15-extSize : l-15]
	e.Size = binary.LittleEndian.Uint32(data[l-11 : l-7])
	e.Block = binary.LittleEndian.Uint32(data[l-15 : l-10])
	return e, nil
}

func (e *EncryptFileInfo) ReadFile(fp string) (*EncryptFileInfo, error) {
	logger.Debug(fp)
	f, err := os.OpenFile(fp, os.O_RDONLY, os.FileMode(0644))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer f.Close()
	buf := make([]byte, 55)
	f.Seek(-55, io.SeekEnd)
	_, err = f.Read(buf)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return e.Read(buf)
}
