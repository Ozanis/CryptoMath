package vector

//Number is an abstraction layer over different vector types that represent big numbers
type Number interface {
	GetCoeficients() []int8
	GetPos(idx uint8) int8
	Sub(vector Vector) Vector
	Add(vector Vector) Vector
}
