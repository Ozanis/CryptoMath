package numeric

func Euler(n uint64) uint64 {
	var res uint64
	if n&1 == 1 {
		res = 2
	} else {
		res = 1
	}
	var p uint64
	for p = 3; p <= n; p += 2 {
		if n%p != 0 {
			if Gcd(n, p) == 1 {
				res++
			}
		}
	}
	return res
}

func Iverse(n, N uint64) uint64 {

	return n
}

func UnityRoot(n, N uint64) uint64 {

	return n
}
