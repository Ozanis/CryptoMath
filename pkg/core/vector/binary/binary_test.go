package binary

import (
	"fmt"
	"testing"
)

var binaryTestCases = []uint8{
	0, 1, 2, 3,
	4, 5, 6, 7,
	8, 9, 10, 11,
	121, 128, 200, 255,
}

func TestRoudDownBin(t *testing.T) {
	expected := []uint8{
		0, 1, 2, 2,
		4, 4, 4, 4,
		8, 8, 8, 8,
		64, 128, 128, 128,
	}
	for i, x := range binaryTestCases {
		y := RoudDownBin(x)
		fmt.Printf("%d: closest bin power of %d is %d\n", i, x, y)
		if y != expected[i] {
			t.Fatalf("Got %d -- want %d\n", y, expected[i])
		}
	}
}

func TestRoudUpBin(t *testing.T) {
	expected := []uint8{
		0, 2, 4, 4,
		8, 8, 8, 8,
		16, 16, 16, 16,
		128, 0, 0, 0,
	}
	for i, x := range binaryTestCases {
		y := RoudUpBin(x)
		fmt.Printf("%d: round up %d to %d\n", i, x, y)
		if y != expected[i] {
			t.Fatalf("Got %d -- want %d\n", y, expected[i])
		}
	}
}

func TestLogBin(t *testing.T) {
	expected := []uint8{
		0, 0, 1, 1,
		2, 2, 2, 2,
		3, 3, 3, 3,
		6, 7, 7, 7,
	}
	for i, x := range binaryTestCases {
		y := LogBin(x)
		fmt.Printf("%d: %d\n", i, y)
		if y != expected[i] {
			t.Fatalf("For log(%b) got %d -- want %d\n", x, y, expected[i])
		}
	}
}

func TestBoolToInt(t *testing.T) {
	test := map[bool]uint8{
		false: 0,
		true:  1,
	}
	for x, check := range test {
		y := BoolToInt(x)
		fmt.Printf("%d: %t\n", y, x)
		if y != check {
			t.Fatalf("Got %d -- want %d\n", y, check)
		}
	}
}

func TestInvert(t *testing.T) {
	expected := []uint8{
		0, 1, 1, 3,
		1, 5, 3, 7,
		1, 9, 5, 13,
		79, 1, 19, 255,
	}
	Log2N := []uint8{
		1, 1, 2, 2,
		3, 3, 3, 3,
		4, 4, 4, 4,
		7, 8, 8, 8,
	}
	for i, x := range binaryTestCases {
		y := Invert(x, Log2N[i])
		fmt.Printf("%d : dec(%d) bin(%b) -- %b\n", i, x, x, y)
		if y != expected[i] {
			t.Fatalf("Got %b -- want: %b", y, expected[i])
		}
	}
}
