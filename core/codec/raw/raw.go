package raw

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
