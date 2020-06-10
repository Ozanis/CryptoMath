package operator

//Equals is an operator of equality for veectors
func Equals(x, y []int8) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

//Greater is an operator that defines one vector is greater then another
func Greater(x, y []int8) bool {
	if len(x) > len(y) {
		return true
	}
	if len(x) < len(y) {
		return false
	}
	for i := range x {
		if x[i] > y[i] {
			return true
		}
	}
	return false
}

//GreaterEquals is an operator that defines one vector is greater or equals then another
func GreaterEquals(x, y []int8) bool {
	if len(x) < len(y) {
		return false
	}
	if len(x) > len(y) {
		return true
	}
	for i := range x {
		if x[i] >= y[i] {
			return true
		}
		return false
	}
	return true
}

//Lesser is an operator that defines one vector is lesser then another
func Lesser(x, y []int8) bool {
	if len(x) < len(y) {
		return true
	}
	if len(x) > len(y) {
		return false
	}
	for i := range x {
		if x[i] < y[i] {
			return true
		}
		return false
	}
	return false
}

//LesserEquals is an operator that defines one vector is lesser or equals then another
func LesserEquals(x, y []int8) bool {
	if len(x) > len(y) {
		return false
	}
	for i := range x {
		if x[i] > y[i] {
			return false
		}
	}
	return true
}
