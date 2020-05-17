package vector

//Len is a warepper arround the builtin len()-1 which works for Numbers
func Len(vector Number) uint8 {
	return uint8(len(vector.GetCoeficients())) - 1
}

//Equals is an operator of equality for veectors
func Equals(vector1, vector2 Number) bool {
	if Len(vector1) != Len(vector2) {
		return false
	}
	for i := Zero; i < Len(vector1); i++ {
		if vector1.GetPos(i) != vector2.GetPos(i) {
			return false
		}
	}
	return true
}

//Greater is an operator that defines one vector is greater then another
func Greater(vector1, vector2 Number) bool {
	if Len(vector1) > Len(vector2) {
		return true
	}
	for i := Zero; i < Len(vector1); i++ {
		if vector1.GetPos(i) > vector2.GetPos(i) {
			return true
		}
	}
	return false
}

//GreaterEquals is an operator that defines one vector is greater or equals then another
func GreaterEquals(vector1, vector2 Number) bool {
	if Len(vector1) < Len(vector2) {
		return false
	}
	for i := Zero; i < Len(vector1); i++ {
		if vector1.GetPos(i) >= vector2.GetPos(i) {
			return true
		}
	}
	return true
}

//Lesser is an operator that defines one vector is lesser then another
func Lesser(vector1, vector2 Number) bool {
	if Len(vector1) < Len(vector2) {
		return true
	}
	for i := Zero; i < Len(vector1); i++ {
		if vector1.GetPos(i) > vector2.GetPos(i) {
			return false
		}
	}
	return true
}

//LesserEquals is an operator that defines one vector is lesser or equals then another
func LesserEquals(vector1, vector2 Number) bool {
	if Len(vector1) > Len(vector2) {
		return false
	}
	for i := Zero; i < Len(vector1); i++ {
		if vector1.GetPos(i) > vector2.GetPos(i) {
			return false
		}
	}
	return true
}
