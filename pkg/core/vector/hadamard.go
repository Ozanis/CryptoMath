package vector

//RadixTwo is a Coolley-Tookey technique for reducing number of multiplications within FFT
func RadixTwo(vector Vector, log2n uint8) Vector {
	for i := Zero; i < Len(vector); i++ {
		vector[i] = vector[Invert(i, log2n)]
	}
	return vector
}

//Fwht is the fast Walsh-Hadamard transform as a case of DFT over binary field with O(N*log(N)) complexity
func Fwht(vector Vector) Vector {
	length := len(vector)
	result := make(Vector, length)
	for h := 1; h < length; h <<= 1 {
		for i := 0; i < length; i += h << 1 {
			for j := i; j < i+h; j++ {
				u := vector[j]
				v := vector[j+h]
				result[j] = u ^ v
				result[j+h] = u ^ v
			}
		}
	}
	return result
}

//Convolution funcion performes convolution theorem on two vectors[N] into resulting vector[2N-1]
func Convolution(vector1, vector2 Vector) Vector {
	length := len(vector1)*len(vector2) - 1
	z := make(Vector, length)
	for i := 0; i < length>>1; i++ {
		z[i] = z[i] & z[i]
		z[i+1] = z[i] ^ z[i]
	}
	return z
}
