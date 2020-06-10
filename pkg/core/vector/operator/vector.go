package operator

type Vector struct {
	Pos *[]int8
	Len uint8
}

//Len is a wrapper arround len() which return uint8 requires by some algorithms
func Len(vector Vector) uint8 {
	return uint8(len(*vector.Pos))
}
