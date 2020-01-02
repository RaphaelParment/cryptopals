package conversion

import "errors"

// Xor computes the xor of input against comparator
func Xor(input byte, comparator byte) byte {
	var output byte
	inputBits := UByteToBits(input)
	comparatorBits := UByteToBits(comparator)
	
	for i := 0; i < 8; i++ {
		if inputBits[i] != comparatorBits[i] {
			output += 1 << (7 - i%8)
		}
	}
	return output
}

// FixedXor computes the input XORed comparator
func FixedXor(input []byte, comparator []byte) ([]byte, error) {
	var output []byte
	if !(len(input) == len(comparator)) {
		return nil, errors.New("length error")
	}

	for i := 0; i < len(input); i++ {
		output = append(output, Xor(input[i], comparator[i]))
	}
	return output, nil
}
