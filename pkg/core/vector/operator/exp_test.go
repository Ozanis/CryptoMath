package operator

import (
	"reflect"
	"testing"
)

var (
	expTestCasesX = [][]int8{
		{1, 1, 0, 1},
	}
	expTestCasesY = [][]int8{
		{0, 1, 1, 0},
	}
)

func TestExp(t *testing.T) {
	expected := [][]int{
		{1, 0, 0, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1},
	}
	for i, want := range expected {
		got := Exp(polynomTestCasesX[i], modTestCase[i])
		if reflect.DeepEqual(got, want) {
			t.Fatal(i, polynomTestCasesX[i], " ^ ", polynomTestCasesY[i], " expected: ", want, " -- got : ", got)
		}
	}
}
