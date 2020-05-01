package ukernel

type Bin32u uint32

const (
	Zero32u = Bin32u(0)
	One32u  = Bin32u(1)
)

var (
	intToSign = []int8{
		0: 1,
		1: -1,
	}

	lookupTable = []Bin32u{
		0b0:                                0,
		0b1:                                0,
		0b10:                               1,
		0b100:                              2,
		0b1000:                             3,
		0b10000:                            4,
		0b100000:                           5,
		0b1000000:                          6,
		0b10000000:                         7,
		0b100000000:                        8,
		0b1000000000:                       9,
		0b10000000000:                      10,
		0b100000000000:                     11,
		0b1000000000000:                    12,
		0b10000000000000:                   13,
		0b100000000000000:                  14,
		0b1000000000000000:                 15,
		0b10000000000000000:                16,
		0b100000000000000000:               17,
		0b1000000000000000000:              18,
		0b10000000000000000000:             19,
		0b100000000000000000000:            20,
		0b1000000000000000000000:           21,
		0b10000000000000000000000:          22,
		0b100000000000000000000000:         23,
		0b1000000000000000000000000:        24,
		0b10000000000000000000000000:       25,
		0b100000000000000000000000000:      26,
		0b1000000000000000000000000000:     27,
		0b10000000000000000000000000000:    28,
		0b100000000000000000000000000000:   29,
		0b1000000000000000000000000000000:  30,
		0b10000000000000000000000000000000: 31,
	}
)

func (B Bin32u) MaxBinPower() Bin32u {
	B |= B >> 1
	B |= B >> 2
	B |= B >> 4
	B |= B >> 8
	B |= B >> 16
	return B >> 1
}

func (B Bin32u) LogBin() Bin32u {
	return lookupTable[B.MaxBinPower()]
}

func (B Bin32u) BinToSign() int8 {
	return intToSign[B]
}

func (B Bin32u) Transpose() Bin32u {
	B = ((B & 0xaaaaaaaa) >> 1) | ((B & 0x55555555) << 1)
	B = ((B & 0xcccccccc) >> 2) | ((B & 0x33333333) << 2)
	B = ((B & 0x0f0f0f0f0) >> 4) | ((B & 0x0f0f0f0f) << 4)
	B = ((B & 0xff00ff00) >> 8) | ((B & 0x00ff00ff) << 8)
	return (B >> 16) | (B << 16)
}

func (B Bin32u) Reverse(log2n Bin32u) Bin32u {
	n := Zero32u
	for i := Zero32u; i < log2n; i++ {
		n <<= 1
		n |= B & 1
		B >>= 1
	}
	return n
}

func BoolToInt(x bool) int8 {
	switch x {
	case true:
		return 1
	default:
		return 0
	}
}
