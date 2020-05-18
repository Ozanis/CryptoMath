package vector

import (
	"reflect"
	"testing"
)

var (
	polynomTestCasesX = [][]int8{
		{1, 1, 0, 1},
		{1, 1, 0, 1, 1, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 0, 1, 1, 0, 1},
	}
	polynomTestCasesY = [][]int8{
		{0, 1, 1, 0},
		{1, 1, 0, 1, 1, 0},
		{1, 0, 1, 0, 1, 0, 1, 0},
		{1, 1, 0, 1, 1, 0},
	}
	modTestCase = [][]int8{
		{1, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0},
	}
)

func TestSub(t *testing.T) {
	expected := [][]int8{
		{1, 1, 0},
		{0, 0, 0, 0, 0, 0},
		{1, 0, 1, 0, 1, 1, 0},
		{0, 1, 0, 1, 1, 1},
	}
	operator := NativeVector{}
	for i, want := range expected {
		got := operator.Sub(polynomTestCasesX[i], polynomTestCasesY[i], modTestCase[i])
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
	operator := NativeVector{}
	for i, want := range expected {
		got := operator.Add(polynomTestCasesX[i], polynomTestCasesY[i], modTestCase[i])
		if reflect.DeepEqual(got, want) {
			t.Fatal(i, polynomTestCasesX[i], " + ", polynomTestCasesY[i], " expected: ", want, " -- got : ", got)
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
	operator := NativeVector{}
	for i, want := range expected {
		got := operator.Mul(polynomTestCasesX[i], (polynomTestCasesY[i]), modTestCase[i])
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
	operator := NativeVector{}
	for i, want := range expected {
		got := operator.Sqr(polynomTestCasesX[i], modTestCase[i])
		if reflect.DeepEqual(got, want) {
			t.Fatal(i, " squared ", polynomTestCasesX[i], " expected: ", want, " -- got : ", got)
		}
	}
}

func TestExp(t *testing.T) {
	expected := [][]int{
		{1, 0, 0, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1},
	}
	operator := NativeVector{}
	for i, want := range expected {
		got := operator.Exp(polynomTestCasesX[i], modTestCase[i])
		if reflect.DeepEqual(got, want) {
			t.Fatal(i, polynomTestCasesX[i], " ^ ", polynomTestCasesY[i], " expected: ", want, " -- got : ", got)
		}
	}
}

func TestInv(t *testing.T) {

}

func TestMod(t *testing.T) {
	expected := [][]int8{
		{0, 1, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 0, 1, 0},
		{0, 1, 1, 0, 1, 1},
	}
	operator := NativeVector{}
	for i, want := range expected {
		got := operator.Mod(polynomTestCasesX[i], polynomTestCasesY[i])
		if reflect.DeepEqual(got, want) {
			t.Fatal(i, polynomTestCasesX[i], " % ", polynomTestCasesY[i], " expected: ", want, " -- got : ", got)
		}
	}
}

func TestDiv(t *testing.T) {
	expected := [][]int8{}
	operator := NativeVector{}
	for i, want := range expected {
		got := operator.Div(polynomTestCasesX[i], polynomTestCasesY[i], modTestCase[i])
		if reflect.DeepEqual(got, want) {
			t.Fatal(i, polynomTestCasesX[i], " / ", polynomTestCasesY[i], " expected: ", want, " -- got : ", got)
		}
	}
}
