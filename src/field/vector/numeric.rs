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

fn ExtendedEuclidian() -> {}

fn LegendreSymbol() -> {}

fn Euler() -> {}

fn IsPrime() -> {}

*/
