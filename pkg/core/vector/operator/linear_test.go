package operator

import (
	"reflect"
	"testing"
)

var (
	polynomTestCasesX = []Vector{
		{1, 1, 0, 1},
		{1, 1, 0, 1, 1, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 0, 1, 1, 0, 1},
	}
	polynomTestCasesY = []Vector{
		{0, 1, 1, 0},
		{1, 1, 0, 1, 1, 0},
		{1, 0, 1, 0, 1, 0, 1, 0},
		{1, 1, 0, 1, 1, 0},
	}
)

func TestSub(t *testing.T) {
	expected := []Vector{
		{1, 1, 0},
		{0, 0, 0, 0, 0, 0},
		{1, 0, 1, 0, 1, 1, 0},
		{0, 1, 0, 1, 1, 1},
	}
	for i, want := range expected {
		got := Sub(polynomTestCasesX[i], polynomTestCasesY[i])
		if reflect.DeepEqual(got, want) {
			t.Fatal(i, polynomTestCasesX[i], " - ", polynomTestCasesY[i], " expected: ", want, " -- got : ", got)
		}
	}
}

func TestAdd(t *testing.T) {
	expected := [][]int8{
		{1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 0},
		{1, 0, 0, 0, 1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0, 1, 1},
	}
	for i, want := range expected {
		got := Add(polynomTestCasesX[i], polynomTestCasesY[i])
		if reflect.DeepEqual(got, want) {
			t.Fatal(i, polynomTestCasesX[i], " + ", polynomTestCasesY[i], " expected: ", want, " -- got : ", got)
		}
	}
}

func TestDiv(t *testing.T) {
	expected := [][]int8{}
	for i, want := range expected {
		got := Div(polynomTestCasesX[i], polynomTestCasesY[i])
		if reflect.DeepEqual(got, want) {
			t.Fatal(i, polynomTestCasesX[i], " / ", polynomTestCasesY[i], " expected: ", want, " -- got : ", got)
		}
	}
}
