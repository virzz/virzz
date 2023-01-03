package basex

// RFC 4648

func basePadding(s string, bit int) string {
	n := bit - len(s)%bit
	if n == bit {
		return s
	}
	for n > 0 {
		n--
		s += "="
	}
	return s
}
