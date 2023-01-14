package utils

import (
	"fmt"
	"testing"
)

func TestCalcPermutation(t *testing.T) {
	var i int
	for i = 1; i <= 11; i++ {
		sum := CalcPermutation(11, i)
		if sum < 0 {
			panic("CalcPermutation overflow")
		}
		fmt.Printf("i = %d sum = %d\n", i, sum)

	}
}
