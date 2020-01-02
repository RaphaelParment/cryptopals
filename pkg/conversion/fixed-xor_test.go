package conversion

import (
	"encoding/hex"
	"testing"
)

func TestFixedXor(t *testing.T) {

	tt := []struct {
		name       string
		input      string
		comparator string
		output     string
	}{{
		name:       "cryptopals",
		input:      "1c0111001f010100061a024b53535009181c",
		comparator: "686974207468652062756c6c277320657965",
		output:     "746865206b696420646f6e277420706c6179",
	}}

	for _, test := range tt {
		in, _ := hex.DecodeString(test.input)
		cmp, _ := hex.DecodeString(test.comparator)
		out, _ := hex.DecodeString(test.output)

		output, err := FixedXor(in, cmp)
		if err != nil {
			t.Fatalf("error: %v", err)
		}

		for i := 0; i < len(in); i++ {
			if output[i] != out[i] {
				t.Fatalf("output byte %d expected %d, got %d", i, out[i], output[i])
			}
		}
	}
}
