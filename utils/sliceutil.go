package utils

import (
	"math/rand"
	"strconv"
)

// PruneEmptyStrings from the slice
func PruneEmptyStrings(v []string) []string {
	return PruneEqual(v, "")
}

// PruneEqual removes items from the slice equal to the specified value
func PruneEqual[T comparable](inputSlice []T, equalTo T) (r []T) {
	for i := range inputSlice {
		if inputSlice[i] != equalTo {
			r = append(r, inputSlice[i])
		}
	}
	return
}

// Dedupe removes duplicates from a slice of elements preserving the order
func Dedupe[T comparable](inputSlice []T) (result []T) {
	seen := make(map[T]struct{})
	for _, inputValue := range inputSlice {
		if _, ok := seen[inputValue]; !ok {
			seen[inputValue] = struct{}{}
			result = append(result, inputValue)
		}
	}
	return
}

// PickRandom item from a slice of elements
func PickRandom[T any](v []T) T {
	return v[rand.Intn(len(v))]
}

// SliceContains if a slice contains an element
func SliceContains[T comparable](inputSlice []T, element T) bool {
	for _, inputValue := range inputSlice {
		if inputValue == element {
			return true
		}
	}
	return false
}

func Intersection[T comparable](s1, s2 []T) []T {
	var r []T
	for _, e := range s1 {
		if SliceContains(s2, e) {
			r = append(r, e)
		}
	}
	return r
}

// ContainsItems checks if s1 contains s2
func ContainsItems[T comparable](s1 []T, s2 []T) bool {
	for _, e := range s2 {
		if !SliceContains(s1, e) {
			return false
		}
	}
	return true
}

// ToInt converts a slice of strings to a slice of ints
func SliceToInt(s []string) ([]int, error) {
	var ns []int
	for _, ss := range s {
		n, err := strconv.Atoi(ss)
		if err != nil {
			return nil, err
		}
		ns = append(ns, n)
	}
	return ns, nil
}
