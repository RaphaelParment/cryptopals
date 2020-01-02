package conversion

import "errors"

// FixedXor computes the input XORed comparator
func FixedXor(input []byte, comparator []byte) ([]byte, error) {
	var (
		output []byte
		v      byte
		used   uint8
	)
	if !(len(input) == len(comparator)) {
		return nil, errors.New("length error")
	}

	inputBits := make([]byte, len(input)*8)
	comparatorBits := make([]byte, len(comparator)*8)

	for i := 0; i < len(input); i++ {
		bits := UByteToBits(input[i])
		for j := 0; j < 8; j++ {
			inputBits[j+(i*8)] = bits[j]
		}

		bits = UByteToBits(comparator[i])
		for j := 0; j < 8; j++ {
			comparatorBits[j+(i*8)] = bits[j]
		}
	}

	for i := 0; i < len(inputBits); i++ {
		if !(inputBits[i] == comparatorBits[i]) {
			v += 1 << (7 - i%8)
		}
		used++
		if used == 8 {
			output = append(output, v)
			v = 0
			used = 0
		}
	}

	return output, nil
}
