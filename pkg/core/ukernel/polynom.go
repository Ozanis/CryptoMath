package ukernel

/*
type Polynom struct {
	Vector []int8
	Len    Bin32u
}

func (P Polynom) Mul(X, Mod Polynom, Domain Fourier) Polynom {
	return P
}

func (P Polynom) Sqr(Mod Polynom, Domain Fourier) Polynom {
	return P
}

func (P Polynom) Pow(X, Mod Polynom, Domain Fourier) Polynom {
	return P
}

func (P Polynom) Add(X Polynom) Polynom {
	for i := Zero32u; i <= X.Len; i++ {
		P.Vector[i] += X.Vector[i]
		P.Vector[i+1] += BoolToInt(X.Vector[i] >= 10)
		P.Vector[i] %= 10
	}
	return P
}

func (P Polynom) Sub(X Polynom) Polynom {
	var carry int8
	for i := Zero32u; i <= X.Len; i++ {
		P.Vector[i] -= X.Vector[i]
		carry = BoolToInt(P.Vector[i] < 0)
		P.Vector[i] = (P.Vector[i] * carry) % 10
		P.Vector[i+1] -= carry
	}
	return P
}

func (P Polynom) Inv(Mod Polynom, Domain Fourier) Polynom {
	return P.Pow(
		Mod.Sub(
			Polynom{
				Vector: []int8{2},
				Len:    1,
			}), Mod, Domain)
}

func (P Polynom) InvPow(X, Mod Polynom, Domain Fourier) Polynom {
	return P.Pow(
		Mod.Sub(
			Polynom{
				Vector: []int8{2},
				Len:    1,
			}).Sub(X), Mod, Domain)
}

func (P Polynom) Div() Polynom {
	return P
}

func (P Polynom) Mod(X Polynom) Polynom {
	if X.Len > P.Len {
		return P
	}
	for i := Zero32u; i <= X.Len; i++ {
		P.Vector[i] &= ^X.Vector[i]
	}
	return P
}

func (P Polynom) Cmp(X Polynom) (Polynom, Polynom) {

	return P, X
}
*/
