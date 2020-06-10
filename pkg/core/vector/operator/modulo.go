package operator

//Mod is a vector(polynomial) modulus operator
func Mod(dest, x, mod Vector) Vector {
	for i := range mod {
		dest.Pos[i] = x.Pos[i] ^ mod.Pos[i]
	}
	return dest
}
