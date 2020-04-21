package api

const Zero32 = uint32(0)

var intBool = map[bool]int8{false: 0, true: 1}
var signBool = map[int8]int8{0: 1, 1: -1}

type Native struct {
	Vector []int8
	Len    uint32
}

func (N Native) Mul(X, Domain, Mod Native) Native {
	N.dft(Domain, Mod)
	X.dft(Domain, Mod)
	for i := Zero32; i <= X.Len*N.Len-1; i++ {
		N.Vector[i] *= X.Vector[i]
		N.Vector[i+1] += (N.Vector[i] / 10) * intBool[N.Vector[i] >= 10]
		N.Vector[i] %= 10
	}
	return N.idft(Domain, Mod)
}

func (N Native) Sqr(Domain, Mod Native) Native {
	N.dft(Domain, Mod)
	for i := Zero32; i <= N.Len; i++ {
		N.Vector[i] *= N.Vector[i]
		N.Vector[i+1] += (N.Vector[i] / 10) * intBool[N.Vector[i] >= 10]
	}
	return N.idft(Domain, Mod)
}

func (N Native) Pow(X, Domain, Mod Native) Native {
	return X
}

func (N Native) Add(X Native) Native {
	for i := Zero32; i <= X.Len; i++ {
		N.Vector[i] += X.Vector[i]
		N.Vector[i+1] += intBool[X.Vector[i] >= 10]
		N.Vector[i] %= 10
	}
	return N
}

func (N Native) Sub(X Native) Native {
	var carry int8
	for i := Zero32; i <= X.Len; i++ {
		N.Vector[i] -= X.Vector[i]
		carry = intBool[N.Vector[i] < 0]
		N.Vector[i] = (N.Vector[i] * signBool[carry]) % 10
		N.Vector[i+1] -= carry
	}
	return N
}

func (N *Native) Inv(Domain, Mod Native) Native {
	return N.Pow(
		Mod.Sub(
			Native{
				Vector: []int8{2},
				Len:    1,
			}), Domain, Mod)
}

func (N *Native) InvPow(X, Domain, Mod Native) Native {
	return N.Pow(
		Mod.Sub(
			Native{
				Vector: []int8{2},
				Len:    1,
			}).Sub(X), Domain, Mod)
}

func (N *Native) Mod(X Native) Native {
	if X.Len > N.Len {
		return *N
	}
	for i := Zero32; i <= X.Len; i++ {
		N.Vector[i] &= ^X.Vector[i]
	}
	return *N
}

func (N Native) Cmp(X Native) (Native, Native) {
	if N.Len > X.Len {
		return N, X
	}
	if X.Len > N.Len {
		return X, N
	}
	for i := Zero32; i < X.Len; i++ {
		if N.Vector[i] > X.Vector[i] {
			return N, X
		} else if X.Vector[i] > N.Vector[i] {
			return X, N
		} else if N.Vector[i] == X.Vector[i] {
			break
		}
	}
	return N, X
}

func (N Native) dft(Domain, mod Native) Native {
	if N.Len == 1 {
		return N
	}
	left := N.odd().dft(Domain.odd(), mod)
	right := N.even().dft(Domain.odd(), mod)
	for i := Zero32; i <= N.Len; i++ {
		root := right.Vector[i] * Domain.Vector[i]
		N.Vector[i] = left.Vector[i] + root
		N.Vector[i+left.Len] = left.Vector[i] - root
	}
	return N

}

func (N Native) idft(Domain, Mod Native) Native {
	N.dft(Domain, Mod)
	I := N.Inv(Domain, Mod)
	for i := Zero32; i < N.Len; i++ {
		N.Vector[i] = N.Vector[i] * I.Vector[i]
		N.Vector[i+1] += (N.Vector[i] / 10) * intBool[N.Vector[i] >= 10]
	}
	return N
}

func (N *Native) split(idx uint32) Native {
	newLen := N.Len >> 2
	newNative := Native{
		Vector: make([]int8, newLen),
		Len:    newLen,
	}
	for i := idx; i <= newLen; i += 2 {
		newNative.Vector[i] = N.Vector[i]
	}
	return newNative
}

func (N Native) odd() Native {
	return N.split(1)
}

func (N Native) even() Native {
	return N.split(0)
}

func (N Native) IsZero() bool {
	var zero bool
	for i := Zero32; i <= N.Len; i++ {
		zero = N.Vector[i] != 0
		if zero == false {
			return false
		}
	}
	return true
}
