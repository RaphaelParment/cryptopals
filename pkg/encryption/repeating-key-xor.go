package encryption

import (
	c "raphael.parment/cryptopals/pkg/conversion"
)

// RepeatingKeyXor func
func RepeatingKeyXor(input, key []byte) []byte {
	var output []byte
	for i := range input {
		xored := c.Xor(input[i], key[i%len(key)])
		output = append(output, xored)
	}

	return output
}
