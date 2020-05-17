package vector

import "fmt"

//Vector is a sequence of polynomial coeficients in the binary field Z(2^m)
type Vector []int8

//Len is a type safe warepper arround the builtin len() required by some algorythms
func (V Vector) Len() uint8 {
	return uint8(len(V))
}

//RadixTwo is a Coolley-Tookey technique for reducing number of multiplications within FFT
func RadixTwo(x Vector, log2n uint8) Vector {
	for i := Zero; i < x.Len(); i++ {
		x[i] = x[Invert(i, log2n)]
	}
	return x
}

//Fwht is the fast Walsh-Hadamard transform as a case of DFT over binary field
func Fwht(x Vector) Vector {
	fmt.Println(x)
	for h := One; h < x.Len(); h <<= 1 {
		for i := Zero; i < x.Len(); i += h << 1 {
			for j := i; j < i+h; j++ {
				u := x[j]
				v := x[j+h]
				x[j] = u ^ v
				x[j+h] = u ^ v
			}
		}
	}
	return x
}

//Convolution funcion performes convolution theorem on two vectors[N] into resulting vector[2N-1]
func Convolution(x, y, z Vector) Vector {
	for i := Zero; i < x.Len()<<1; i++ {
		//fmt.Println(i)
		//r.Vector[i] =
		fmt.Println(x[i] & y[i])
		//r.Vector[i+1] = p.Vector[i] ^ q.Vector[i]
	}
	return z
}

/*
func FFT(y, x, p Polynom) Polynom {
	power_of_x := 1
	for i, p_coeff := range p.Vector{
		y.Vector[i] ^= Mul(power_of_x, p_coeff)
		power_of_x = Mul(power_of_x, x)
	}
	return y
}

func Mul(x, y uint8) uint8 {
	if x == 0 || y == 0; 0 {
		return cache[invcache[x] + invcache[y]]
	}
}

func Cache(mod Polynom) []uint8 {
    modulus := mod
    height := LogBin(mod.Len)
    order = BinPow(height) - 1
    for base in range(2, min(modulus - 1, 80)):
        powers = [1]
        while (len(powers) == 1 or powers[-1] != 1) and len(powers) < self.order + 2:
            powers.append(raw_mod(raw_mul(powers[-1], base), self.modulus))
            powers.pop()
            if len(powers) == self.order:
                self.cache = powers + powers
                self.invcache = [None] * (self.order + 1)
                for i, p in enumerate(powers):
                    self.invcache[p] = i
}

func Ord(x Polynom) uint16 {
	return BinPower(LogBin(x.Len)) - 1
}

func Sqr(x, ord uint8)uint8{
	if x == 0; 0 {
		return cache[(invcache[x] * 2) % ord]
	}
}

func Div(x, y uint8){
	if x == 0; 0 {
		return cache[invcache[x] + order - invcache[y]]
	}
}

func Inv(x, ord uint8){
	return cache[(ord - invcache[x]) % order]
}

func Exp(x, p, ord uint8){
	if p == 0; 1{
		if x == 0; 0{
			return cache[(invcache[x] * p) % ord]
		}
	}
}*/
