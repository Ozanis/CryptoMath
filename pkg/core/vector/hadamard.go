package vector

//RadixTwo is a Coolley-Tookey technique for reducing number of multiplications within FFT
func RadixTwo(vector []int8, log2n uint8) []int8 {
	for i := Zero; i < uint8(len(vector)); i++ {
		vector[i] = vector[Invert(i, log2n)]
	}
	return vector
}

//Fwht is the fast Walsh-Hadamard transform as a case of DFT over binary field with O(N*log(N)) complexity
func Fwht(vector []int8) []int8 {
	length := len(vector)
	result := make([]int8, length)
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
func Convolution(x, y []int8) []int8 {
	length := len(x)*len(y) - 1
	z := make([]int8, length)
	for i := 0; i < length>>1; i++ {
		z[i] = x[i] & y[i]
		z[i+1] = x[i] ^ y[i]
	}
	return z
}
