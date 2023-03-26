package encrypt

import (
	"io"
	"os"
	"path"
	"testing"

	"github.com/virzz/logger"
)

func TestEncrypt(t *testing.T) {
}

func TestAesCTREncrypt(t *testing.T) {
	src := []byte("asdfgh1234567890asdfgh1234567890asdfgh1234567890asdfgh1234567890asdfgh1234567890asdfgh1234567890asdfgh1234567890asdfgh1234567890asdfgh1234567890asdfgh1234567890asdfgh1234567890asdfgh1234567890asdfgh1234567890asdfgh1234567890asdfgh1234567890asdfgh1234567890")
	dst, err := AesCTREncrypt(src, []byte("asdfgh1234567890asdfgh1234567890"), []byte("1234567890asdfgh"))
	if err != nil {
		t.Fatal(err)
	}
	logger.Success(len(src))
	logger.Success(len(dst))
}

func TestEncryptFileInfo(t *testing.T) {
	e := &EncryptFileInfo{}
	e.Ext = []byte("png")
	e.Size = 234
	e.Block = 764423
	e.IsCompress = true
	n, d := e.Data()
	logger.Info(n, string(d), d)
	e2 := &EncryptFileInfo{}
	e2.Read(d)
	logger.Info(e2.Block, e2.Size, string(e2.Ext))
}

func TestExt(t *testing.T) {
	logger.Info(path.Ext("xxx.png"))
	logger.Info(path.Base("xxx.png"))
	logger.Info(path.Ext("xxx.png.zst"))
	logger.Info(path.Base("xxx.png.zst"))
}

func TestT(t *testing.T) {
	fn := path.Join("/tmp", "test_trun")
	f, err := os.OpenFile(fn, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	b := make([]byte, 2048)
	for i := 0; i < 2048; i++ {
		b[i] = 'a'
	}
	f.Write(b)
	f.Seek(0, 0)
	b = make([]byte, 1024)
	for i := 0; i < 1024; i++ {
		b[i] = 'b'
	}
	// f.Seek(512, 0)
	f.Seek(0, 0)
	f.WriteAt(b, 0)
	b = []byte("tessssstttttttttttttttttttttttt")
	f.Seek(0, io.SeekEnd)
	f.Write(b)
	// f.Truncate(512 + 1024)
}
