package binary

import "unsafe"

const (
	//Zero predefined constant for uint zero representation
	Zero = uint8(0)
	//One predefined constant for uint one's representation
	One = uint8(1)
)

//RoudDownBin returns 2 base number rounded down in O(log(N))
func RoudDownBin(x uint8) uint8 {
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	return x - (x >> 1)
}

//RoudUpBin is derivative function to round up number to near power of 2 in O(log(N))
func RoudUpBin(x uint8) uint8 {
	return RoudDownBin(x) << 1
}

//IsBinRadix returns 1 in case if given number is radix of 2 and 0 otherwise in O(1)
func IsBinRadix(x uint8) uint8 {
	return (x & (x - 1)) ^ 0
}

//Invert implements bit inversion with O(log(N)) required by FFT
func Invert(x, n uint8) uint8 {
	var y uint8
	for i := Zero; i < n; i++ {
		y <<= 1
		y |= x & 1
		x >>= 1
	}
	return y
}

//LogBin finds log with  base 2 of 32 bit number with O(log(N)) complexity and without branching
func LogBin(x uint8) uint8 {
	y := BoolToInt(x > 0xF) << 2
	x >>= y
	shift := BoolToInt(x > 0x3) << 1
	x >>= shift
	y |= shift
	return x>>1 | y
}

//BoolToInt used for fast conversion bool to uint8
func BoolToInt(x bool) uint8 {
	return *(*uint8)(unsafe.Pointer(&x)) & 1
}

//IntToBool used for fast conversion from uint16 to bool
func IntToBool(x uint8) bool {
	x &= 1
	return *(*bool)(unsafe.Pointer(&x))
}
