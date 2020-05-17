package vector

//Number is an abstraction layer over different vector types that represent big numbers
type Number interface {
	GetCoeficients() []int8
	GetPos(idx uint8) int8
	Sub(vector Vector) Vector
	Add(vector Vector) Vector
	Sqr() Vector
	Mul(vector Vector) Vector
	Exp(vector Vector) Vector
	Inv(vector Vector) Vector
	Mod(vector Vector) Vector
	Div(vector Vector) Vector
}
