package models

import (
	"fmt"
)

const (
	pRebinding = "r_"
)

// GetRebinding -
func GetRebinding(key string) (result string, err error) {
	return Get(fmt.Sprintf("%s_%s", pRebinding, key))
}

// SetRebinding -
func SetRebinding(key, value string) (result string, err error) {
	return Set(fmt.Sprintf("%s_%s", pRebinding, key), value)
}
