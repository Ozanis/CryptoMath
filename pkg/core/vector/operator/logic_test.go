package operator

import "testing"

var (
	compareTestCasesX = [][]int8{
		{1, 1, 0, 1},
		{1, 1, 0, 1, 1, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 0, 1, 1, 0, 1},
	}
	compareTestCasesY = [][]int8{
		{0, 1, 1, 0},
		{1, 1, 0, 1, 1, 0},
		{1, 0, 1, 0, 1, 0, 1, 0},
		{1, 1, 0, 1, 1, 0},
	}
)

func TestEquals(t *testing.T) {
	expected := []bool{false, true, false, false}
	for i, got := range expected {
		got = Equals(compareTestCasesX[i], compareTestCasesY[i])
		if got != expected[i] {
			t.Fatal(i, " expected: ", expected[i], " -- got : ", got)
		}
	}
}

func TestGreater(t *testing.T) {
	expected := []bool{true, false, true, true}
	for i, got := range expected {
		got = Greater(compareTestCasesX[i], compareTestCasesY[i])
		if got != expected[i] {
			t.Fatal(i, " expected: ", expected[i], " -- got : ", got)
		}
	}
}

func TestGreaterEquals(t *testing.T) {
	expected := []bool{true, true, true, true}
	for i, got := range expected {
		got = GreaterEquals(compareTestCasesX[i], compareTestCasesY[i])
		if got != expected[i] {
			t.Fatal(i, " expected: ", expected[i], " -- got : ", got)
		}
	}
}

func TestLesser(t *testing.T) {
	expected := []bool{false, false, false, false}
	for i, got := range expected {
		got = Lesser(compareTestCasesX[i], compareTestCasesY[i])
		if got != expected[i] {
			t.Fatal(i, " expected: ", expected[i], " -- got : ", got)
		}
	}
}

func TestLesserEquals(t *testing.T) {
	expected := []bool{false, true, false, false}
	for i, got := range expected {
		got = LesserEquals(compareTestCasesX[i], compareTestCasesY[i])
		if got != expected[i] {
			t.Fatal(i, " expected: ", expected[i], "-- got : ", got)
		}
	}
}
