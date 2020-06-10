package operator

/*
//Euclidian uses extend algorithm to find multiplicative inverse
func Euclidian(x, mod []int8) []int8 {
	var q,t []int8
	m := mod
	y := 0
	v := 1;
	if Equals(mod, x){
		return 0;
	}
	for a > 1; {
		q = Div(x, mod)
		t = mod;
		m = Mod(x, mod)
		x = t;
		t = y;
		y = Sub(v, Mul(q, y))
		v = t;
	}
	if v < 0{
	   v = Add(v, m)
	}
	return v;
	}
}


func Gcd(x, y []int8) []int8{
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

func Legendre() []int8 {}
*/
