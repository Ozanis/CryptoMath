package numeric

func Gcd(a, b uint64) uint64 {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a == 1 || b == 1 {
		return 1
	}
	var shift uint64
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
