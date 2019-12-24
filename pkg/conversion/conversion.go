package conversion

// UByteToBits converts unsigned byte <b> to a
// bit array of uint8
func UByteToBits(b uint8) []uint8 {
	bits := make([]uint8, 8)
	cmp := uint8(1)
	for i := 0; i < 8; i++ {
		if b&cmp == 1 {
			bits[7-i] = 1
		} else {
			bits[7-i] = 0
		}
		b = b >> 1
	}

	return bits
}
