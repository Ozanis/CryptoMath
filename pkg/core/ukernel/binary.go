package ukernel

import "unsafe"

const (
	//Zero32u is predefined constant of 0 for iteration with 32 bit counter
	Zero32u = uint32(0)
	//One32u is predefined constant of 1 for iteration with 32 bit counter
	One32u = uint32(1)
)

//Predifined slice used for calculation carryes without branching
var (
	IntToSign = []int8{
		0: 1,
		1: -1,
	}
)

//BinPower returns 2 base number rounded down in O(log(N))
func BinPower(x uint32) uint32 {
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	return x - (x >> 1)
}

//NextBinPower is derivative function to round up number to near power of 2 in O(log(N))
func NextBinPower(x uint32) uint32 {
	return BinPower(x) << 1
}

//IsBinRadix returns 1 in case if given number is radix of 2 and 0 otherwise in O(1)
func IsBinRadix(x uint32) uint32 {
	return (x & (x - 1)) ^ Zero32u
}

//Reverse 32 bit number with 5 * parallel log(N) complexity
func Reverse(x uint32) uint32 {
	x = ((x & 0xaaaaaaaa) >> 1) | ((x & 0x55555555) << 1)
	x = ((x & 0xcccccccc) >> 2) | ((x & 0x33333333) << 2)
	x = ((x & 0x0f0f0f0f0) >> 4) | ((x & 0x0f0f0f0f) << 4)
	x = ((x & 0xff00ff00) >> 8) | ((x & 0x00ff00ff) << 8)
	return (x >> 16) | (x << 16)
}

//Transpose implements bit inversion required by FFT
func Transpose(x, n uint32) uint32 {
	var y uint32
	for i := Zero32u; i < n; i++ {
		y <<= 1
		y |= x & 1
		x >>= 1
	}
	return y
}

//LogBin finds log with  base 2 of 32 bit number with log(N) complexity and without branching
func LogBin(x uint32) uint32 {
	y := BoolToInt(x > 0xFFFF) << 4
	x >>= y
	shift := BoolToInt(x > 0xFF) << 3
	x >>= shift
	y |= shift
	shift = BoolToInt(x > 0xF) << 2
	x >>= shift
	y |= shift
	shift = BoolToInt(x > 0x3) << 1
	x >>= shift
	y |= shift
	return x>>1 | y
}

//BoolToInt used for fast conversion bool to uint32
func BoolToInt(x bool) uint32 {
	return *(*uint32)(unsafe.Pointer(&x)) & One32u
}

//IntToBool used for fast conversion from uint32 to bool
func IntToBool(x uint32) bool {
	x &= One32u
	return *(*bool)(unsafe.Pointer(&x))
}
