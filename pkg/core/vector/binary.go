package vector

import "unsafe"

const (
	//Zero16 precalculated 16 bit zero representation
	Zero16 = uint16(0)
	//One16 precalculated 16 bit one's representation
	One16 = uint16(1)
)

//BinPower returns 2 base number rounded down in O(log(N))
func BinPower(x uint16) uint16 {
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	return x - (x >> 1)
}

//NextBinPower is derivative function to round up number to near power of 2 in O(log(N))
func NextBinPower(x uint16) uint16 {
	return BinPower(x) << 1
}

//IsBinRadix returns 1 in case if given number is radix of 2 and 0 otherwise in O(1)
func IsBinRadix(x uint16) uint16 {
	return (x & (x - 1)) ^ 0
}

//Invert implements bit inversion with O(log(N)) required by FFT
func Invert(x, n uint16) uint16 {
	var y uint16
	for i := Zero16; i < n; i++ {
		y <<= 1
		y |= x & 1
		x >>= 1
	}
	return y
}

//LogBin finds log with  base 2 of 32 bit number with O(log(N)) complexity and without branching
func LogBin(x uint16) uint16 {
	y := BoolToInt16(x > 0xFF) << 3
	x >>= y
	shift := BoolToInt16(x > 0xF) << 2
	x >>= shift
	y |= shift
	shift = BoolToInt16(x > 0x3) << 1
	x >>= shift
	y |= shift
	return x>>1 | y
}

//BoolToInt64 used for fast conversion bool to uint64
func BoolToInt64(x bool) uint64 {
	return *(*uint64)(unsafe.Pointer(&x)) & 1
}

//BoolToInt16 used for fast conversion bool to uint16
func BoolToInt16(x bool) uint16 {
	return *(*uint16)(unsafe.Pointer(&x)) & 1
}

//IntToBool used for fast conversion from uint16 to bool
func IntToBool(x uint16) bool {
	x &= 1
	return *(*bool)(unsafe.Pointer(&x))
}
