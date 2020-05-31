package decryption

import (
	"raphael.parment/cryptopals/pkg/conversion"
	c "raphael.parment/cryptopals/pkg/conversion"
	"raphael.parment/cryptopals/pkg/encryption"
)

// HammingDistance func
func HammingDistance(input, comparator []byte) int {
	dist := 0
	for i := range input {
		iBits := c.UByteToBits(input[i])
		cBits := c.UByteToBits(comparator[i])
		for j := 0; j < 8; j++ {
			if iBits[j] != cBits[j] {
				dist++
			}
		}
	}
	return dist
}

// prepareBlocks divides the input into a 2d array of size numBlocks by keysize
func prepareBlocks(keysize int, numBlocks int, input []byte) [][]byte {
	blocks := make([][]byte, numBlocks)
	k := 0
	for i := 0; i < numBlocks; i++ {
		blocks[i] = make([]byte, keysize)
		for j := 0; j < keysize; j++ {
			blocks[i][j] = input[k]
			k++
		}
	}

	return blocks
}

// FindKeySize returns the keysize based on the Hamming Distance of pairs of blocks
// with length ranging from 2 to maxKeySize
func FindKeySize(input []byte, maxKeysize int) int {
	numBlocks := 32
	keysize := 0

	// Initialise to max 32 bit value
	minDist := float32((1 << 32) - 1)

	for i := 2; i < maxKeysize; i++ {
		hammingDist := float32(0)
		dist := float32(0)
		blocks := prepareBlocks(i, numBlocks, input)

		for j := 0; j < numBlocks; j += 2 {
			hammingDist += float32(HammingDistance(blocks[j], blocks[j+1])) / float32(i)
			dist += hammingDist
		}
		normalisedDist := dist / float32(numBlocks)
		if normalisedDist < minDist {
			minDist = normalisedDist
			keysize = i
		}
	}
	return keysize
}

// PrepareInput func
func PrepareInput(input []byte, keysize int) [][]byte {
	k := 0
	noBlocks := len(input) / keysize
	in := make([][]byte, noBlocks)

	// Break input into blocks of keysize bytes
	for i := 0; i < noBlocks; i++ {
		in[i] = make([]byte, keysize)
		for j := 0; j < keysize; j++ {
			in[i][j] = input[k]
			k++
		}
	}

	// Transpose
	transposed := make([][]byte, keysize)
	for i := 0; i < keysize; i++ {
		transposed[i] = make([]byte, noBlocks)
		for j := 0; j < noBlocks; j++ {
			transposed[i][j] = in[j][i]
		}
	}
	return transposed
}

// Solve func
func Solve(input []byte, transposed [][]byte, keysize int) []byte {
	var repeatedKey []byte
	for i := 0; i < keysize; i++ {
		_, key := conversion.FindSingleXorKey(transposed[i])
		repeatedKey = append(repeatedKey, key)
	}
	return encryption.RepeatingKeyXor(input, repeatedKey)
}
