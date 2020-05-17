package vector

import (
	"encoding/binary"
)

const (
	//Mask64 is a precision of the 64 base number
	Mask64 = uint64(0xffffffffffffffff)
)

//Native is implementation of converter to vector with 64 bit precision
type Native struct{}

//MapBytes - bytes to uint64 mapping
func (Native) MapBytes(bytes []byte) interface{} {
	return binary.LittleEndian.Uint64(bytes)
}

//ConvertToBase - convet to 64bit positioning system
func (Native) ConvertToBase(stream []uint64) Vector {
	fromStream := make(Vector, len(stream))
	cache := Mask64
	for i := 0; i <= len(stream); i++ {
		if cache >= 0 {
			cache -= stream[i]
		} else {
			fromStream[i] = 1
			cache = Mask64
		}
	}
	return fromStream
}

//ToBytes converts Vector type to byte stream
func (Native) ToBytes(vector Vector) []byte {
	bytes := make([]byte, len(vector))
	return bytes
}

//ToVector converts byte stream to Vector type that ready to operate
func (Native) ToVector(bytes []byte) Vector {
	newVector := make(Vector, len(bytes))
	return newVector
}
