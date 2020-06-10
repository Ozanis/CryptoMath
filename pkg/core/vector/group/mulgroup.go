package group

import (
	"cryptops/pkg/core/vector/binary"
	"cryptops/pkg/core/vector/operator"
)

//Mul is multiplication vector operator
func Mul(x, y, mod operator.Vector) operator.Vector {
	return operator.RadixTwo((operator.Fwht(operator.Convolution(operator.Fwht(x), operator.Fwht(y)))))
}

//Sqr is a squaring vector operator
func Sqr(x, mod operator.Vector) operator.Vector {
	fwhtVector := operator.Fwht(x)
	return operator.RadixTwo(operator.Fwht(operator.Convolution(fwhtVector, fwhtVector)), binary.LogBin(operator.Len(mod)))
}
