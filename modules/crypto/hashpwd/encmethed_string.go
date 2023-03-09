// Code generated by "stringer -type=EncMethed"; DO NOT EDIT.

package hashpwd

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MySQL-0]
	_ = x[NTLMv1-1]
	_ = x[MD5-2]
	_ = x[MD5x2-3]
	_ = x[MD5x3-4]
	_ = x[SHA1MD5-5]
	_ = x[SHA1-6]
	_ = x[SHA1x2-7]
	_ = x[MD5SHA1-8]
	_ = x[SHA256-9]
	_ = x[MD5SHA256-10]
}

const _EncMethed_name = "MySQLNTLMv1MD5MD5x2MD5x3SHA1MD5SHA1SHA1x2MD5SHA1SHA256MD5SHA256"

var _EncMethed_index = [...]uint8{0, 5, 11, 14, 19, 24, 31, 35, 41, 48, 54, 63}

func (i EncMethed) String() string {
	if i < 0 || i >= EncMethed(len(_EncMethed_index)-1) {
		return "EncMethed(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EncMethed_name[_EncMethed_index[i]:_EncMethed_index[i+1]]
}