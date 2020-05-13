package vector

const (
	//MaskLow32 - bitmask for getting high order half of uint64
	MaskLow32 = 0x000000003fffffff
	//MaskHigh32 - bitmask for getting low order half of uint64
	MaskHigh32 = 0xfffffffc0000000
	//Mask63 add uint64 without overflow 64 bit
	Mask63 = 0x7ffffffffffffff
	//Mask64th get 64th high order bit
	Mask64th = 0x8000000000000000
	//MaskIm - mask for identifying imaginarity of the number
	MaskIm = 0xff00
	//MaskRe - mask for identifying sign of the number
	MaskRe = 0x00ff
)

//Ariphmetic64 helper type for handling 64bit number products
type Ariphmetic64 struct {
	Weight uint64
	Mask   uint16
}

//Ariphmetic128 helper type for handling 64bit number products
type Ariphmetic128 struct {
	Major, Minor uint64
	Mask         uint16
}

//Add128 - vectorized complex addition
func Add128(x, y Ariphmetic128) Ariphmetic128 {
	x.Minor = Add64(x.Minor, y.Minor)
	x.Major = Add64(x.Major, y.Major)
	x.Major += CarryAdd64(x.Minor, y.Minor)
	return x
}

//Sub128 - vectorized complex addition
func Sub128(x, y Ariphmetic128) Ariphmetic128 {
	x.Minor = Sub64(x.Minor, y.Minor)
	x.Major = Sub64(x.Major, y.Major)
	x.Major -= BoolToInt64(x.Minor < y.Minor)
	return x
}

//Sub64 substraction of uint64 with carrying
func Sub64(x, y uint64) uint64 {
	return (Mask64th | x) - (Mask63 & y)
}

//Add64 addition of uint64 with carrying
func Add64(x, y uint64) uint64 {
	return (Mask63 & x) + (Mask63 & y)
}

//CarryAdd64 calculation of addition carry
func CarryAdd64(x, y uint64) uint64 {
	return ((x >> 60) & (y >> 60)) & ((x >> 61) | (y >> 61))
}

//Prod128 - production of Ariphmetic64 numbers
func Prod128(x, y Ariphmetic64) Ariphmetic128 {
	return Ariphmetic128{
		Major: MajorBits32(y.Weight) * MajorBits32(x.Weight),
		Minor: MinorBits32(x.Weight) * MinorBits32(y.Weight),
		Mask:  x.Mask ^ y.Mask,
	}
}

//MajorBits32 - getting high order part of uint64 number
func MajorBits32(x uint64) uint64 {
	return (x & MaskHigh32) >> 31
}

//MinorBits32 - getting low order of part uint64
func MinorBits32(x uint64) uint64 {
	return x & MaskLow32
}
