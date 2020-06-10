package operator

modTestCase = [][]int8{
	{1, 0, 0, 0, 0, 0, 0, 0},
	{1, 0, 0, 0, 0, 0, 0, 0},
	{1, 0, 0, 0, 0, 0, 0, 0},
	{1, 0, 0, 0, 0, 0, 0, 0},
}

func TestMod(t *testing.T) {
	expected := [][]int8{
		{0, 1, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 0, 1, 0},
		{0, 1, 1, 0, 1, 1},
	}
	testVector := Vector{
		Vector: make([]int8, len(mod)),
		Len:
	}
	for i, want := range expected {
		got := Mod(polynomTestCasesX[i], polynomTestCasesY[i])
		if reflect.DeepEqual(got, want) {
			t.Fatal(i, polynomTestCasesX[i], " % ", polynomTestCasesY[i], " expected: ", want, " -- got : ", got)
		}
	}
}
