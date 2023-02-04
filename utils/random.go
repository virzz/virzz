//go:build !go1.20
// +build !go1.20

package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
