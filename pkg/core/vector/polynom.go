package vector

//NativeVector is realisation of operators using native compiler optimisation
type NativeVector struct{}

//Sub is substraction vector operator
func (NativeVector) Sub(x, y, mod []int8) []int8 {
	var carry int8
	vector := make([]int8, len(mod))
	for i := len(mod); i >= 0; i-- {
		vector[i] = x[i] ^ y[i]
		carry = int8(BoolToInt(x[i] > y[i]))
		vector[i+1] ^= carry
	}
	return x
}

//Add is addition vector operator
func (NativeVector) Add(x, y, mod []int8) []int8 {
	var carry int8
	vector := make([]int8, len(mod))
	for i := len(mod); i >= 0; i-- {
		vector[i] = x[i] ^ y[i] ^ carry
		carry = x[i] & y[i]
	}
	return vector
}

//Mul is multiplication vector operator
func (NativeVector) Mul(x, y, mod []int8) []int8 {
	return RadixTwo((Fwht(Convolution(Fwht(x), Fwht(y)))), LogBin(Len(mod)))
}

//Sqr is a squaring vector operator
func (NativeVector) Sqr(x, mod []int8) []int8 {
	fwhtVector := Fwht(x)
	return RadixTwo(Fwht(Convolution(fwhtVector, fwhtVector)), LogBin(Len(mod)))
}

//Exp is an exponentiation vector operator
func (NativeVector) Exp(x, mod []int8) []int8 {
	return x
}

//Inv finds multiplicative inverse
func (NativeVector) Inv(x, mod []int8) []int8 {
	return x
}

//Mod is a vector(polynomial) modulus operator
func (NativeVector) Mod(x, mod []int8) []int8 {
	modulo := make([]int8, len(mod))
	for i := range modulo {
		modulo[i] = x[i] ^ mod[i]
	}
	return modulo
}

//Div is a vector(polynomial) division operator
func (NativeVector) Div(x, y, mod []int8) []int8 {
	return x
}

//IsIrreducible checks does polynom irreducible
func (NativeVector) IsIrreducible() bool {
	return true
}
