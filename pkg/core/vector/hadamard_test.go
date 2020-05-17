package vector

import (
	"testing"
)

var (
	vectorTestCasesX = []Vector{
		{1, 1, 0, 1},
		{1, 1, 1, 1, 0, 0, 0, 0},
		//{1, 1, 0, 1, 1, 0, 1},
	}
	vectorTestCasesY = []Vector{
		{0, 1, 1, 0},
		{1, 0, 1, 0, 1, 0, 1, 0},
		//{1, 1, 0, 1, 1, 0},
	}
)

func TestRadixTwo(t *testing.T) {

}

func TestFhwt(t *testing.T) {
	expected := []Vector{
		{1, 0, 1, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
	}
	for i := range expected {
		got := Fwht(vectorTestCasesX[i])
		if !Equals(got, expected[i]) {
			t.Fatal(i, " Provided: ", vectorTestCasesX[i], "-- got : ", got, " expected: ", expected[i])
		}
	}
}

func TestConvolution(t *testing.T) {
	expected := []Vector{
		{1, 0, 0, 1, 1, 1, 0},
	}
	for i := 0; i < len(expected)/2; i += 2 {
		result := Fwht(Convolution(Fwht(vectorTestCasesX[i]), Fwht(vectorTestCasesY[i])))
		if !Equals(result, expected[i]) {
			t.Fatal(i, " expected: ", expected[i], "-- got : ", result)
		}
	}
}
