package number

var intToSign = []int8{
	0: 1,
	1: -1,
}

func IntToSign(x int8) int8 {
	return intToSign[x]
}

func BoolToInt(x bool) int8 {
	switch x {
	case true:
		return 1
	default:
		return 0
	}
}
