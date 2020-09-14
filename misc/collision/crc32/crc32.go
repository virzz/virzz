package crc32

import (
	"archive/zip"
	"fmt"
	"hash/crc32"
	"sort"
	"strings"
	"sync"

	"github.com/virink/virzz/misc/collision"
)

// ZipCRC32 -
func ZipCRC32(name, table string, limit ...int) (string, error) {
	length := 4
	if len(limit) > 0 && limit[0] < 7 {
		length = limit[0]
	}
	zipReader, err := zip.OpenReader(name)
	if err != nil {
		return "", err
	}
	defer zipReader.Close()
	res := make(map[string]string, len(zipReader.File))
	var wg sync.WaitGroup
	for _, f := range zipReader.File {
		if f.UncompressedSize64 <= uint64(length) {
			wg.Add(1)
			go bruteForceCRC32(&wg, f.Name, table, f.CRC32, res)
		}
	}
	wg.Wait()

	keys := make([]string, 0)
	for k := range res {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	result := make([]string, len(res))
	for i, k := range keys {
		result[i] = fmt.Sprintf("%s - %s", k, res[k])
	}
	return strings.Join(result, "\n"), nil
}

func bruteForceCRC32(wg *sync.WaitGroup, name, table string, crc uint32, res map[string]string) {
	defer wg.Done()
	dt := collision.NewDictTable()
	dt.SetTable([]byte(table))
	dt.SetLength(4)
	dt.SetCollisionByte(func(secret []byte) bool {
		return crc32.ChecksumIEEE(secret) == crc
	})
	dt.ProcessCollision()
	for _, v := range dt.Results() {
		res[name] = v
		break
	}
}
