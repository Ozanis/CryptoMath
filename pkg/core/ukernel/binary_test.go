package ukernel

import (
	"fmt"
	"testing"
)

var cases = []uint32{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 121, 128, 200, 256,
	4095, 4096, 1073741823, 1073741824, 1073741825, 10066176,
}

func TestBinPower(t *testing.T) {
	expected := []uint32{
		0, 1, 2, 2, 4, 4, 4, 4, 8, 8, 8, 8, 64, 128, 128, 256,
		2048, 4096, 536870912, 1073741824, 1073741824, 8388608,
	}
	for i, x := range cases {
		y := BinPower(x)
		fmt.Printf("%d: closest bin power of %d is %d\n", i, x, y)
		if y != expected[i] {
			t.Fatalf("Got %d -- want %d\n", y, expected[i])
		}
	}
}

func TestNextBinPower(t *testing.T) {
	expected := []uint32{
		0, 2, 4, 4, 8, 8, 8, 8, 16, 16, 16, 16, 128, 256, 256, 512,
		4096, 8192, 1073741824, 2147483648, 2147483648, 16777216,
	}
	for i, x := range cases {
		y := NextBinPower(x)
		fmt.Printf("%d: round up %d to %d\n", i, x, y)
		if y != expected[i] {
			t.Fatalf("Got %d -- want %d\n", y, expected[i])
		}
	}
}

func TestLogBin(t *testing.T) {
	expected := []uint32{
		0, 0, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 6, 7, 7, 8,
		11, 12, 29, 30, 30, 23,
	}
	for i, x := range cases {
		y := LogBin(x)
		fmt.Printf("%d: %d\n", i, y)
		if y != expected[i] {
			t.Fatalf("For log(%b) got %d -- want %d\n", x, y, expected[i])
		}
	}
}

func TestReverse(t *testing.T) {
	expected := []uint32{
		0, 2147483648, 1073741824, 3221225472, 536870912, 2684354560, 1610612736, 3758096384,
		268435456, 2415919104, 1342177280, 3489660928, 2650800128, 16777216, 318767104,
		8388608, 4293918720, 524288, 4294967292, 2, 2147483650, 10066176,
	}
	for i, x := range cases {
		y := Reverse(x)
		fmt.Printf("%d : dec(%d) bin(%b) -- %b\n", i, x, x, y)
		if y != expected[i] {
			t.Fatalf("Got %b -- want: %b", y, expected[i])
		}
	}
}

func TestBoolToInt(t *testing.T) {
	test := map[bool]uint32{
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

func TestTranspose(t *testing.T) {
	expected := []uint32{
		0, 1, 1, 3, 1, 5, 3, 7, 1, 9, 5, 13, 79, 1, 19, 1,
		4095, 1, 1073741823, 1, 1073741825, 39321,
	}
	Log2N := []uint32{
		1, 1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 7, 8, 8, 9,
		12, 13, 30, 31, 31, 24,
	}
	for i, x := range cases {
		y := Transpose(x, Log2N[i])
		fmt.Printf("%d : dec(%d) bin(%b) -- %b\n", i, x, x, y)
		if y != expected[i] {
			t.Fatalf("Got %b -- want: %b", y, expected[i])
		}
	}
}
