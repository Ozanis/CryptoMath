package bitwise

func NextPowerOfTwo(v uint32) uint32 {
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	return v
}
