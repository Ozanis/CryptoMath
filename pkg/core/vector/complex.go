package vector

//Root is a vector of unity roots in the complex field for implmenting Fourier
type Root struct {
	Sign int8
	Im   uint8
}

//IntToSign - Predefined slice used for calculation carryes without branching
var IntToSign = []int8{
	0: 1,
	1: -1,
}

//Exp is a function for fast unity roots calculation: e^(j*2*Pi*n/p)
func (R Root) Exp(Phi, N uint16) Root {
	return Root{
		Sign: sign(Phi>>2, N),
		Im:   imaginary(Phi>>1, N),
	}
}

//sign - ensure unity root belongs both positive or negative quarter of the complex plane
func sign(Phi, N uint16) int8 {
	return IntToSign[((Phi>>2)/N)&1]
}

//imaginary - ensure unity root belongs both imaginary or real quarter of the complex plane
func imaginary(Phi, N uint16) uint8 {
	return uint8((N >> 1) / Phi)
}
