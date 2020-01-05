package encryption

import (
	"encoding/hex"
	"testing"
)

func TestRepeatingKeyXor(t *testing.T) {
	tt := []struct {
		name   string
		input  string
		key    string
		output string
	}{{
		name:   "cryptopals",
		input:  "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal",
		key:    "ICE",
		output: "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f",
	}}

	for _, test := range tt {
		output := RepeatingKeyXor([]byte(test.input), []byte(test.key))
		out := hex.EncodeToString(output)
		if out != test.output {
			t.Fatalf("failed test '%s', expected: %s, got: %s",
				test.name, test.output, out)
		}
	}
}
