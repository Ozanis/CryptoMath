package converter

//Converter interface is abstraction layer over different vector types
type Converter interface {
	ToBytes(vector []int8) []byte
	ToVector(bytes []byte) []int8
	MapBytes(bytes []byte) []uint64
	ConvertToBase(vector []uint64) []int8
}
