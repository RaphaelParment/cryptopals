package encryption

// RepeatingKeyXor func
func RepeatingKeyXor(input, key []byte) []byte {
	var output []byte
	for i := range input {
		xored := input[i] ^ key[i%len(key)]
		output = append(output, xored)
	}

	return output
}
