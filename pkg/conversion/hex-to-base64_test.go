package conversion

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestHexToBase64(t *testing.T) {

	tt := []struct {
		name   string
		input  string
		output string
	}{{
		name:   "cryptopals",
		input:  "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
		output: "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t",
	}, {
		name:   "Some random test text",
		input:  "536f6d652072616e646f6d20746573742074657874",
		output: "U29tZSByYW5kb20gdGVzdCB0ZXh0",
	}}

	for _, test := range tt {
		in, _ := hex.DecodeString(test.input)
		fmt.Println(string(HexToBase64(in)))
		out := string(HexToBase64(in))
		if out != test.output {
			t.Fatalf("failed to convert input for '%s'. Got: '%s', expected: '%s'", test.name, out, test.output)
		}
	}
}
