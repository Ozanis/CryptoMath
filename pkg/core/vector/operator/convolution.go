package operator

import (
	"cryptops/pkg/core/vector/binary"
	"fmt"
)

//RadixTwo is a Coolley-Tookey technique for reducing number of multiplications within FFT
func RadixTwo(x Vector, log2n uint8) Vector {
	for i := binary.Zero; i < x.Len; i++ {
		x.Pos[i] = x.Pos[binary.Invert(i, log2n)]
	}
	return x
}

//Fwht is the fast Walsh-Hadamard transform as a case of DFT over binary field with O(N*log(N)) complexity
func Fwht(x Vector) Vector {
	for h := binary.One; h < x.Len; h <<= 1 {
		for i := binary.Zero; i < x.Len; i += h {
			for j := i; j < i+h; j++ {
				fmt.Println(i + h)
				u := x.Pos[j]
				v := x.Pos[j+h]
				x.Pos[j] = u ^ v
				x.Pos[j+h] = u ^ v
			}
		}
	}
	return x
}

//Convolution funcion performes convolution theorem on two vectors[N] into resulting vector[2N-1]
func Convolution(dest, x, y Vector) Vector {
	for i := 0; i < dest.Len>>1; i++ {
		dest.Pos[i] = x.Pos[i] & y.Pos[i]
		dest.Pos[i+1] = x.Pos[i] ^ y.Pos[i]
	}
	return dest
}
