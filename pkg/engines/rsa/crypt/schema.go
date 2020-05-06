package crypt

import . "math/rand"

func PublicKey(N int) int {
	return 2*Intn(N/2-1) + 1
}

func PrivateKey(e, phi int) int {
	return Power(e, Euler(phi)-1, phi)
}

func Rsa(msg []int, exp, N int) []int {
	data := make([]int, len(msg))
	for j, i := range msg {
		data[j] = Power(i, exp, N)
	}
	return data
}

func Encode(msg []int) []byte {
	data := make([]byte, len(msg))
	for j, i := range msg {
		data[j] = byte(i)
	}
	return data
}

func Decode(msg []byte) []int {
	data := make([]int, len(msg))
	for j, i := range msg {
		data[j] = int(i)
	}
	return data
}
