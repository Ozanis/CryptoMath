package vector

import (
	"reflect"
	"testing"
)

var VectorTestCase = []Vector{
	{1, 1, 0, 1},
	{0, 1, 1, 0},
	{1, 1, 1, 1, 0, 0, 0, 0},
	{1, 0, 1, 0, 1, 0, 1, 0},
}

func TestFhwt(t *testing.T) {
	expected := []Vector{
		{1, 1, 1, 1},
		{0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	for i := range expected {
		result := Fwht(VectorTestCase[i])
		if !reflect.DeepEqual(result, expected[i]) {
			t.Fatal("expected: ", expected[i], "-- got : ", result)
		}
	}
}

func TestConvolution(t *testing.T) {
	expected := []Vector{
		{1, 0, 0, 1, 1, 1, 0},
	}
	r := Vector{0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < len(expected)/2; i += 2 {
		result := Fwht(Convolution(Fwht(VectorTestCase[i]), Fwht(VectorTestCase[i+1]), r))
		if !reflect.DeepEqual(result, expected) {
			t.Fatal("expected: ", expected, "-- got : ", result)
		}
	}
}
