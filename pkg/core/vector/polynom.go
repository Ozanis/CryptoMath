package vector

//Vector is a sequence of polynomial coeficients in the binary field Z(2^m)
type Vector []int8

//GetCoeficients represents vector as int8 array
func (V Vector) GetCoeficients() []int8 {
	return []int8(V)
}

//GetPos returns coeficint on exact position
func (V Vector) GetPos(idx uint8) int8 {
	return V[idx]
}

//Sub is substraction vector operator
func (V Vector) Sub(vector Vector) Vector {
	carry := int8(0)
	for i, x := range vector {
		V[i] ^= x ^ carry
		carry = V[i] & x
	}
	return V
}

//Add is addition vector operator
func (V Vector) Add(vector Vector) Vector {
	carry := int8(0)
	for i, x := range vector {
		V[i] ^= x ^ carry
		carry = V[i] & x
	}
	return V
}

//Mul is multiplication vector operator
func (V Vector) Mul(vector Vector) Vector {
	return Fwht(Convolution(Fwht(vector), Fwht(V)))
}

//Sqr is a squaring vector operator
func (V Vector) Sqr() Vector {
	fwhtVector := Fwht(V)
	return Fwht(Convolution(fwhtVector, fwhtVector))
}

//Exp is an exponentiation vector operator
func (V Vector) Exp(vector Vector) Vector {
	fwhtVector := Fwht(vector)
	return Fwht(Convolution(fwhtVector, fwhtVector))
}

//Inv finds multiplicative inverse
func (V Vector) Inv(vector Vector) Vector {
	return vector
}

//Mod is a vector(polynomial) modulus operator
func (V Vector) Mod(vector Vector) Vector {
	return vector
}

//Div is a vector(polynomial) division operator
func (V Vector) Div(vector Vector) Vector {
	return vector
}
