package vector

//Converter interface is abstraction layer over different vector types
type Converter interface {
	ToBytes(vector interface{}) []byte
	ToVector(bytes []byte) Vector
	MapBytes(bytes []byte) interface{}
	ConvertToBase(vector interface{}) interface{}
}
