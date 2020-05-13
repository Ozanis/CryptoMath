package vector

//Complex pair of polynoms both consists of real and imaginary parts separately
type Complex struct {
	Vector   []Ariphmetic128
	Carry    uint64
	N, Log2N uint16
}

//Butterfly relises indexes permutation for reducing number of multiplications
func Butterfly(complex Complex) Complex {
	complex.Log2N = LogBin(complex.N)
	var inv uint16
	for i := Zero16; i < F.Log2N; i++ {
		complex.Vector[i] = complex.Vector[Invert(i, complex.Log2N)]
	}
	return complex
}

//Dft is the core function for implementing convolution theorem
func Dft(complex Complex) Fourier {
	for s := One16; s <= complex.Log2N; s++ {
		m := One16 << s            // 2 power s
		m2 := m >> One16           // m2 = m/2 -1
		wm := UnityRoot(calculating.Vector[], m2) // principle root of nth complex - root of unity.
		for j := Zero16; j < m2; j++ {
			for k := j; k < F.N; k += m {
				t := F.UnityRoot(k + m2) // t = twiddle factor
				u := F.Vector[k]
				F.Vector[k] = Add(u, t)    // calculating y[k]
				F.Vector[k+m2] = Sub(u, t) // similar calculating y[k+n/2]
			}
			w.Mul(wm)
		}
	}
	return F
}

//UnityRoot is a function for fast unity roots calculation: w * e^(j*2*Pi*n/p)
func UnityRoot(x Ariphmetic128, phi uint16) Ariphmetic128 {
	x.Mask ^= re(Phi, x.N)
	x.Mask ^= im(Phi, x.N)
	return x
}

//re - ensure real part sign
func re(Phi, N uint16) uint16 {
	return (((Phi << 2) / N) & 1) * MaskRe
}

//im - ensure imaginarity
func im(Phi, N uint16) uint16 {
	return (Phi << 1) / N * MaskIm
}


/*
func Idft(polynom, Domain []Complex) Polynom {
	N := uint32(len(polynom))
	idft := Number{
		Vector: make([]int8, N),
		Len:    N,
	}
	for i := Zero32; i <= N - 1; i++ {
		idft.Vector[i] = sumReal(mulComplex(polynom[i], Domain[i]))
	}
	idft.Re
}


func (F Fourier) Dft(Domain Fourier) Fourier {
	Y := F.butterfly()
	for s := One32u; s <= F.Log2N; s++ {
		m := One32u << s  // 2 power s
		m2 := m >> One32u // m2 = m/2 -1
		for j := Zero32u; j < m2; j++ {
			for k := j; k < F.N; k += m {
				t := Domain.Vector[m2].Mul(F.Vector[k+m2]) // t = twiddle factor
				u := Y.Vector[k]
				Y.Vector[k] = u.Add(t)    // calculating y[k]
				Y.Vector[k+m2] = u.Sub(t) // similar calculating y[k+n/2]
			}
		}
	}
	return Y
}


func (F Fourier) IdftStateless(polynom, Domain []Complex) Number {
	N := uint32(len(polynom))
	idft := Number{
		Vector: make([]int8, N),
		Len:    N,
	}
	for i := Zero32; i <= N - 1; i++ {
		idft.Vector[i] = sumReal(mulComplex(polynom[i], Domain[i]))
	}
}
*/
