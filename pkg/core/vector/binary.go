package vector

import "unsafe"

const (
	//MaskLow32 - bitmask for getting high order half of number
	MaskLow32 = int64(0x0000ffff)
	//MaskHigh32 - bitmask for getting low order half of number
	MaskHigh32 = int64(0xffff0000)
	//Zero16 - predefined 16 bit representation of zero
	Zero16 = uint16(0)
	//One16 - predefined 16 bit representation of one
	One16 = uint16(1)
)

//HighHalf - getting high order part of int64 number
func highBits32(x int64) int64 {
	return (x & MaskHigh32) >> 32
}

//LowHalf - getting low order of part int64
func lowBits32(x int64) int64 {
	return x & MaskLow32
}

//HalfMul - multiplication for avoid 64bit overflow
func HalfMul(x, y int64) (int64, int64) {
	return lowBits32(x) * lowBits32(y), highBits32(y) * highBits32(x)
}

//BinPower returns 2 base number rounded down in O(log(N))
func BinPower(x uint16) uint16 {
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	//x |= x >> 16
	return x - (x >> 1)
}

//NextBinPower is derivative function to round up number to near power of 2 in O(log(N))
func NextBinPower(x uint16) uint16 {
	return BinPower(x) << 1
}

//IsBinRadix returns 1 in case if given number is radix of 2 and 0 otherwise in O(1)
func IsBinRadix(x uint16) uint16 {
	return (x & (x - 1)) ^ Zero16
}

//Transpose implements bit inversion required by FFT
func Transpose(x, n uint16) uint16 {
	var y uint16
	for i := Zero16; i < n; i++ {
		y <<= 1
		y |= x & 1
		x >>= 1
	}
	return y
}

//LogBin finds log with  base 2 of 32 bit number with log(N) complexity and without branching
func LogBin(x uint16) uint16 {
	y := BoolToInt(x > 0xFF) << 3
	x >>= y
	shift := BoolToInt(x > 0xF) << 2
	x >>= shift
	y |= shift
	shift = BoolToInt(x > 0x3) << 1
	x >>= shift
	y |= shift
	return x>>1 | y
}

//BoolToInt used for fast conversion bool to uint32
func BoolToInt(x bool) uint16 {
	return *(*uint16)(unsafe.Pointer(&x)) & One16
}

//IntToBool used for fast conversion from uint32 to bool
func IntToBool(x uint16) bool {
	x &= One16
	return *(*bool)(unsafe.Pointer(&x))
}
