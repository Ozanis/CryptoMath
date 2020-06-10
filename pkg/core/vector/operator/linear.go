package operator

import (
	"cryptops/pkg/core/vector/binary"
)

//Sub is substraction vector operator
func Sub(dest, x, y Vector) Vector {
	var carry int8
	for i := dest.Len; i >= 0; i-- {
		dest.Pos[i] = x.Pos[i] ^ y.Pos[i]
		carry = int8(binary.BoolToInt(x.Pos[i] > y.Pos[i]))
		dest.Pos[i+1] -= carry
	}
	return dest
}

//Add is addition vector operator
func Add(dest, x, y Vector) Vector {
	var carry uint8
	for i := mod.Len; i >= 0; i-- {
		dest.Pos[i] = x.Pos[i] ^ y.Pos[i] ^ carry
		carry = x.Pos[i] & y.Pos[i]
	}
	return dest
}

//Div is a vector(polynomial) division operator without remainder
func Div(dest, x, y Vector) Vector {

	return x
}
