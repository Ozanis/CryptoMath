package crypt

func Power(base, exp, mod int) int {
	if (mod == 1) || (base == 0) {
		return 0
	}
	if (exp == 0) || (base == 1) {
		return 1
	}
	res := Power(base, exp/2, mod)
	res = (res * res) % mod
	if exp&1 != 0 {
		return ((base % mod) * res) % mod
	}
	return res % mod
}

func Gcd(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a == 1 || b == 1 {
		return 1
	}
	var shift int
	for shift = 1; ((a | b) & 1) == 0; shift <<= 1 {
		a >>= 1
		b >>= 1
	}
	for a > 0 {
		if (a & 1) == 0 {
			a >>= 1
		}
		if (b & 1) == 0 {
			b >>= 1
		}
		if a < b {
			t := b
			b = a
			a = t
		}
		a -= b
	}
	return b * shift
}

func EulerMul(p, q int) int {
	return (p - 1) * (q - 1)
}

func Euler(x int) int {
	var phi int
	if x&1 == 1 {
		phi = 2
	} else {
		phi = 1
	}
	var p int
	for p = 3; p <= x; p += 2 {
		if x%p != 0 {
			if Gcd(x, p) == 1 {
				phi++
			}
		}
	}
	return phi
}
