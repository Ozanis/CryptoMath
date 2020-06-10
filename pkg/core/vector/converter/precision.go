package converter

import (
	"encoding/binary"
)

const (
	//Mask64 is a precision of the 64 base number
	Mask64 = uint64(0xffffffffffffffff)
)

//Precision is implementation of converter to vector with 64 bit precision
type Precision struct{}

//MapBytes - bytes to uint64 mapping
func (Precision) MapBytes(bytes []byte) uint64 {
	return binary.LittleEndian.Uint64(bytes)
}

//ConvertToBase - convet to 64bit positioning system
func (Precision) ConvertToBase(stream []uint64) []int8 {
	vector := make([]int8, len(stream))
	cache := Mask64
	for i := 0; i <= len(stream); i++ {
		if cache >= 0 {
			cache -= stream[i]
		} else {
			vector[i] = 1
			cache = Mask64
		}
	}
	return vector
}

//ToBytes converts Vector type to byte stream
func (Precision) ToBytes(vector []int8) []byte {
	bytes := make([]byte, len(vector))
	return bytes
}

//ToVector converts byte stream to Vector type that ready to operate
func (Precision) ToVector(bytes []byte) []int8 {
	newVector := make([]int8, len(bytes))
	return newVector
}
