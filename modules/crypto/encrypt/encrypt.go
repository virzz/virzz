package encrypt

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
	"github.com/virzz/logger"
	"github.com/virzz/virzz/utils/execext"
)

const (
	EncryptBlockSize int64 = 1024 * 10
)

func Compress(filepath string, unCompress bool) error {
	cmd := "zstd -9 %s"
	if unCompress {
		if !strings.HasSuffix(filepath, ".zst") {
			return errors.New("compressed file must with ext .zst")
		}
		cmd = `zstd -f -d %s`
	}
	cmd = fmt.Sprintf(cmd, path.Base(filepath))
	// var stdout bytes.Buffer
	opts := &execext.RunCommandOptions{
		Command: cmd,
		Dir:     path.Dir(filepath),
		// Stdout:  &stdout,
		// Stderr:  os.Stderr,
	}
	if err := execext.RunCommand(context.Background(), opts); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func AesCTREncrypt(data, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	dst := make([]byte, len(data))
	cipher.NewCTR(block, iv).XORKeyStream(dst, data)
	return dst, nil
}

func Decrypt(filename string, key, iv []byte, efi *EncryptFileInfo) error {
	f, err := os.OpenFile(filename, os.O_RDWR, os.FileMode(0644))
	if err != nil {
		return err
	}
	defer f.Close()
	if len(iv) == 0 {
		buf := make([]byte, aes.BlockSize)
		_, err = f.ReadAt(buf, int64(efi.Size)-aes.BlockSize)
		if err != nil && err != io.EOF {
			return errors.WithStack(err)
		}
		iv = buf
	}
	f.Seek(0, 0)
	buffer := make([]byte, efi.Block)
	n, err := io.ReadFull(f, buffer)
	if err != nil && err != io.EOF {
		return errors.WithStack(err)
	}
	logger.Debug("Decrypt")
	decData, err := AesCTREncrypt(buffer[:n], key, iv)
	if err != nil {
		return errors.WithStack(err)
	}
	f.Seek(0, 0)
	f.WriteAt(decData, 0)
	f.Truncate(int64(efi.Size))

	if efi.IsCompress {
		if err := Compress(filename, true); err != nil {
			return errors.WithStack(err)
		}
		logger.SuccessF("%s decrypted and uncompressed, oldfile %s.old", filename, filename)
	} else {
		logger.SuccessF("%s decrypted", filename)
	}
	return nil
}

func Check(filename string) error {
	efi := &EncryptFileInfo{}
	// IsEncrypted
	efi, err := efi.ReadFile(filename)
	if err != nil {
		if errors.Is(err, ErrNotEncryptedByVirzz) {
			logger.Success(err)
			return nil
		}
		return errors.WithStack(err)
	}
	logger.SuccessF(`Encrypt By VIRZZ
	Origin Size : %d
	Block  Size : %d
	IsCompressed: %v
	Ext         : %s`, efi.Size, efi.Block, efi.IsCompress, efi.Ext)
	return nil
}

func Encrypt(filename string, key, iv []byte, compress bool) error {
	oriName := filename
	efi := &EncryptFileInfo{}
	// IsEncrypted
	efi, err := efi.ReadFile(filename)
	if err != nil && !errors.Is(err, ErrNotEncryptedByVirzz) {
		return errors.WithStack(err)
	}
	if efi != nil {
		return Decrypt(filename, key, iv, efi)
	}
	efi = &EncryptFileInfo{}
	efi.Ext = []byte(path.Ext(filename))
	if compress && strings.HasSuffix(filename, ".zst") {
		logger.Warn("It seems to have been compressed")
	} else if compress {
		err := Compress(filename, false)
		if err != nil {
			return errors.WithStack(err)
		} else {
			efi.IsCompress = true
			filename = fmt.Sprintf("%s.zst", filename)
		}
	}
	f, err := os.OpenFile(filename, os.O_RDWR, os.FileMode(0644))
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	stat, err := f.Stat()
	if err != nil {
		return errors.WithStack(err)
	}
	encryptBlockSize := EncryptBlockSize
	size := stat.Size()
	for size < encryptBlockSize+aes.BlockSize && encryptBlockSize > 1024 {
		encryptBlockSize -= 1024
	}
	logger.Debug("encryptBlockSize = ", encryptBlockSize)
	if size < encryptBlockSize+aes.BlockSize {
		return errors.New("File size is too small.")
	}
	efi.Size = uint32(size)
	efi.Block = uint32(encryptBlockSize)
	buffer := make([]byte, EncryptBlockSize)
	n, err := io.ReadFull(f, buffer)
	if err != nil && err != io.EOF {
		return errors.WithStack(err)
	}
	if len(iv) == 0 {
		f.Seek(0, 0)
		buf := make([]byte, aes.BlockSize)
		_, err = f.ReadAt(buf, size-aes.BlockSize)
		if err != nil && err != io.EOF {
			return errors.WithStack(err)
		}
		iv = buf
	}
	logger.Debug("Encrypt")
	encData, err := AesCTREncrypt(buffer[:n], key, iv)
	if err != nil {
		return errors.WithStack(err)
	}
	f.Seek(0, 0)
	_, err = f.Write(encData)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = f.Seek(0, io.SeekEnd)
	if err != nil {
		return errors.WithStack(err)
	}
	_, extData := efi.Data()
	_, err = f.Write(extData)
	if err != nil {
		return errors.WithStack(err)
	}
	if efi.IsCompress {
		logger.SuccessF("%s compressed and encrypted to %s", oriName, filename)
	} else {
		logger.SuccessF("%s encrypted", filename)
	}
	return nil
}
