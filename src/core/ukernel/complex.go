package ukernel

type Complex [2]int8

func (C Complex) Mul(X Complex) Complex {
	for i := 0; i <= 1; i++ {
		for j := 0; j <= 1; j++ {
			C[i] *= X[j]
		}
	}
	return C
}

func (C Complex) Add(X Complex) Complex {
	for i := 0; i <= 1; i++ {
		C[i] += X[i]
	}
	return C
}

func (C Complex) Sub(X Complex) Complex {
	for i := 0; i <= 1; i++ {
		C[i] -= X[i]
	}
	return C
}

func (C Complex) Arg(Phi, N Bin32u) Complex {
	angle := (Phi >> 2) / N
	im := int8((Phi>>1)/N) * (1 & angle).BinToSign()
	re := ^im
	return Complex{
		re,
		im,
	}
}
