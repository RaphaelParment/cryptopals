package decryption

import (
	"crypto/cipher"
	"fmt"
)

// DecryptECD decrypts
func DecryptECD(in []byte, b cipher.Block) ([]byte, error) {
	size := len(in)
	if size%b.BlockSize() != 0 {
		return nil, fmt.Errorf("incorrect block size")
	}

	out := make([]byte, size)
	for i := 0; i < size; i += b.BlockSize() {
		b.Decrypt(out[i:], in[i:])
	}

	return out, nil
}

// DeteDetectECB checks if the input <in> is "probably" ECB encoded
func DetectECB(in []byte, size int) (bool, error) {
	if len(in)%size != 0 {
		return false, fmt.Errorf("incorrect block size")
	}

	seen := make(map[string]struct{})
	for i := 0; i < len(in); i += size {
		val := string(in[i : i+size])
		if _, ok := seen[val]; ok {
			return true, nil
		}
		seen[val] = struct{}{}
	}
	return false, nil
}
