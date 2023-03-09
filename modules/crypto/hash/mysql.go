// Code from https://github.com/go-sql-driver/mysql/blob/master/auth.go

package hash

import (
	"crypto/sha1"
	"fmt"
)

// Generate binary hash from byte string using insecure pre 4.1 method
func MySQLHash(password []byte) string {
	var add, tmp, r1, r2 uint32 = 7, 0, 1345345333, 0x12345671
	for _, c := range password {
		if c == ' ' || c == '\t' {
			continue // skip spaces and tabs in password
		}
		tmp = uint32(c)
		r1 ^= (((r1 & 63) + add) * tmp) + (r1 << 8)
		r2 += (r2 << 8) ^ r1
		add += tmp
	}
	// Remove sign bit (1<<31)-1)
	return fmt.Sprintf("*%08x%08x", r1&0x7FFFFFFF, r2&0x7FFFFFFF)
}

// Generate binary Hash password using 4.1+ method (SHA1)
func MySQL5Hash(password []byte) string {
	has := sha1.Sum(password)
	has = sha1.Sum(has[:])
	return fmt.Sprintf("*%x", has[:])
}
