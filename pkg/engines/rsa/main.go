package main

import (
	"crypt"
	"fmt"
)

const plain = "cryptoanalysis"

func main() {
	msg := []byte(plain)

	p := 17
	q := 29
	N := p * q
	phi := crypt.EulerMul(p, q)

	public := crypt.PublicKey(phi)
	private := crypt.PrivateKey(public, phi)

	cipher := crypt.Rsa(crypt.Decode(msg), public, N)
	open := crypt.Rsa(cipher, private, N)

	fmt.Println("Plain text: ", plain)
	fmt.Println("Encoded: ", msg)
	fmt.Println("Encrypted: ", cipher)
	fmt.Println("Decrypted: ", open)
	fmt.Println("Decoded: ", string(crypt.Encode(open)))

	fmt.Println("p = ", p)
	fmt.Println("q = ", q)
	fmt.Println("p * q = ", N)
	fmt.Println("public key = ", public)
	fmt.Println("private key = ", private)
	fmt.Println("e * d mod N = ", (public*private)%phi)

}
