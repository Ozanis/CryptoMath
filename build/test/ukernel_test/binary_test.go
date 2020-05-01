package ukernel_test

import (
	"core/ukernel"
	"fmt"
	"testing"
)

var cases = []ukernel.Bin32u{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 121, 128, 200, 256, 333, 343, 353, 339, 510, 511, 512, 640, 641, 700, 1000, 1024,
	(2 << 11) - 1, 2 << 11, (2 << 11) + 1, (2 << 20) - 1, 2 << 20, (2 << 20) + 1, (2 << 30) - 1, 2 << 30, (2 << 30) + 1,
}

func TestMaxBinPower(t *testing.T) {
	expected := []ukernel.Bin32u{
		0, 0, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 6, 7, 7, 8, 8, 8, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 10,
		10, 11, 11, 19, 20, 20, 29, 30, 30,
	}

	for i, x := range cases {
		res := x.MaxBinPower()
		fmt.Println("%d: %d -- Expected: %d", i, res, expected[i])
		if res != expected[i] {
			t.Fatal("Got %d -- want %d", res, expected[i])
		}
	}
}

func TestLogBin(t *testing.T) {
	expected := []ukernel.Bin32u{
		0, 0, 2, 2, 4, 4, 4, 4, 8, 8, 8, 8, 64, 128, 128, 256, 256, 256, 256, 256, 256, 256, 512, 512, 512, 512, 512, 1024,
		2 << 10, 2 << 11, 2 << 11, 2 << 19, 2 << 20, 2 << 20, 2 << 29, 2 << 30, 2 << 30,
	}
	for i, x := range cases {
		res := x.LogBin()
		fmt.Println("%d: %d -- Expected: %d", i, res, expected[i])
		if res != expected[i] {
			t.Fatal("Got %d -- want %d", res, expected[i])
		}
	}
}

func TestBitReverse(t *testing.T) {
	expected := []ukernel.Bin32u{
		0b0, 0b1 << 31, 0b10 << 30, 0b11 << 30, 0b100 << 29, 0b101 << 29, 0b110 << 29, 0b111 << 29, 0b1000 << 28, 0b1001 << 28,
		0b1010 << 28, 0b1011 << 28, 0b1111001 << 25, 0b10000000 << 24, 0b11001000 << 24, 0b10000000 << 24, 0b101001101 << 23,
		0b101010111 << 23, 0b101100001 << 23, 0b101010011 << 23, 0b111111110 << 23, 0b111111111 << 23, 0b1000000000 << 22,
		0b1010000000 << 22, 0b1010000001 << 22, 0b1010111100 << 22, 0b1111101000 << 22, 0b10000000000 << 21, 0b11111111111 << 20,
		0b100000000000 << 20, 0b100000000001 << 20, 0b11111111111111111111 << 12, 0b100000000000000000000 << 11,
		0b100000000000000000001 << 11, 0b111111111111111111111111111111 << 2,
		0b10, 0b1000000000000000000000000000001 << 1,
	}
	for i, x := range cases {
		res := x.Transpose()
		fmt.Println("%d: %d -- Expected: %d", i, res, expected[i])
		if res != expected[i] {
			t.Fatal("Got %d -- want %d", res, expected[i])
		}
	}
}
