package operator

import (
	"reflect"
	"testing"
)

var (
	vectorTestCasesX = [][]int8{
		{1, 1, 0, 1},
		{1, 1, 1, 1, 0, 0, 0, 0},
	}
	vectorTestCasesY = [][]int8{
		{0, 1, 1, 0},
		{1, 0, 1, 0, 1, 0, 1, 0},
	}
)

func TestRadixTwo(t *testing.T) {

}

func TestFhwt(t *testing.T) {
	expected := [][]int8{
		{1, 0, 1, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
	}
	for i := range expected {
		got := Fwht(vectorTestCasesX[i])
		if !reflect.DeepEqual(got, expected[i]) {
			t.Fatal(i, " Provided: ", vectorTestCasesX[i], "-- got : ", got, " expected: ", expected[i])
		}
	}
}

func TestConvolution(t *testing.T) {
	expected := [][]int8{
		{1, 0, 0, 1, 1, 1, 0},
	}
	for i := range expected {
		result := Fwht(Convolution(Fwht(vectorTestCasesX[i]), Fwht(vectorTestCasesY[i])))
		if !reflect.DeepEqual(result, expected[i]) {
			t.Fatal(i, " expected: ", expected[i], "-- got : ", result)
		}
	}
}

func TestMul(t *testing.T) {
	expected := [][]int8{
		{1, 0, 0, 1, 1, 1, 0},
		{1, 0, 1, 1, 0, 1, 1, 0, 0, 1, 0, 0},
		{1, 0, 0, 1, 1, 1, 1, 1, 0, 1, 1, 0, 0, 0, 0, 0},
		{1, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0},
	}
	for i, want := range expected {
		got := Mul(polynomTestCasesX[i], (polynomTestCasesY[i]), modTestCase[i])
		if reflect.DeepEqual(got, want) {
			t.Fatal(i, polynomTestCasesX[i], " - ", polynomTestCasesY[i], " expected: ", want, " -- got : ", got)
		}
	}
}

func TestSqr(t *testing.T) {
	expected := [][]int8{
		{1, 0, 1, 0, 1, 0, 0, 1},
		{1, 0, 1, 1, 0, 1, 1, 0, 0, 1, 0, 0},
		{1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 1, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1},
	}
	for i, want := range expected {
		got := Sqr(polynomTestCasesX[i], modTestCase[i])
		if reflect.DeepEqual(got, want) {
			t.Fatal(i, " squared ", polynomTestCasesX[i], " expected: ", want, " -- got : ", got)
		}
	}
}
