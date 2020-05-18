package vector

//Number is an abstraction layer over different vector (polynom) types that represent big numbers
type Number interface {
	Sqr(x []int8) []int8
	Sub(x, y, mod []int8) []int8
	Add(x, y, mod []int8) []int8
	Mul(x, y, mod []int8) []int8
	Exp(x, y, mod []int8) []int8
	Inv(x, y, mod []int8) []int8
	Mod(x, y, mod []int8) []int8
	Div(x, y, mod []int8) []int8
	IsIrreducible(x []int8) bool
}
