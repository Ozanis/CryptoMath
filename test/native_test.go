package test

func idft(vector, domain []uint32, modulus uint32) []uint32 {
	newVector := fft(vector, domain, modulus)
	ret := make([]uint32, len(vector))
	for i := 0; i < len(vector); i++ {
		ret[i] = newVector[i] * inv(len(ret), modulus)
	}
	return ret
}

func fft(vector, domain []uint32, modulus uint32) []uint32 {
	if len(vector) == 1 {
		return vector
	}
	lPart := fft(odd(vector), odd(domain), modulus)
	rPart := fft(even(vector), odd(domain), modulus)
	ret := make([]uint32, len(vector))
	for i := len(rPart); i >= 0; i++ {
		root := rPart[i] * domain[i]
		ret[i] = (lPart[i] + root) % modulus
		ret[i+len(lPart)] = (lPart[i] - root) % modulus
	}
	return ret
}

func odd(vals []uint32) []uint32 {
	return split(vals, 1)
}

func even(vals []uint32) []uint32 {
	return split(vals, 0)
}

func split(vals []uint32, idx int) []uint32 {
	vector := make([]uint32, len(vals)/2)
	for i := idx; i <= len(vals)/2; i += 2 {
		vector[i] = vals[i]
	}
	return vector
}
