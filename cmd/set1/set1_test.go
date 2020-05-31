package set1

import (
	"bufio"
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"raphael.parment/cryptopals/pkg/conversion"
	"raphael.parment/cryptopals/pkg/decryption"
	"raphael.parment/cryptopals/pkg/encryption"
)

func TestChallenge1(t *testing.T) {
	out := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	in := toBytes(t, "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	output := base64.StdEncoding.EncodeToString(in)
	if output != out {
		t.Fatalf("wrong output: got %v expected %v", output, out)
	}
}

func TestChallenge2(t *testing.T) {
	in := toBytes(t, "1c0111001f010100061a024b53535009181c")
	cmp := toBytes(t, "686974207468652062756c6c277320657965")
	out := toBytes(t, "746865206b696420646f6e277420706c6179")

	output, err := conversion.FixedXor(in, cmp)
	if err != nil {
		t.Fatalf("failed to xor; %v", err)
	}

	if !reflect.DeepEqual(output, out) {
		t.Fatalf("wrong output: got %v expected %v", output, out)
	}
}

func TestChallenge3(t *testing.T) {
	var (
		key, v, char    byte
		score, maxScore float32
	)

	in := toBytes(t, "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	for char = 0; char < 128; char++ {
		for i := 0; i < len(in); i++ {
			v = in[i] ^ char
			score += conversion.CorpusEN[v]
		}
		if score > maxScore {
			maxScore = score
			key = char
		}
		score = 0
	}

	for _, c := range in {
		t.Logf("%c", c^key)
	}

}

func TestChallenge4(t *testing.T) {
	in := getInput(t, "./inputs/4.txt")

	var (
		maxScore float32
		key      byte
		iter     int
	)

	for i, bytes := range in {
		s, k := conversion.FindSingleXorKey(bytes)
		if s > maxScore {
			maxScore = s
			key = k
			iter = i
		}
	}

	t.Logf("iteration: %d, key: %d\n", iter, key)

	for _, b := range in[iter] {
		t.Logf("%c", b^key)
	}
}

func TestChallenge5(t *testing.T) {
	in := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	out := toBytes(t, "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")

	output := encryption.RepeatingKeyXor(in, []byte("ICE"))
	if !reflect.DeepEqual(output, out) {
		t.Fatalf("wrong output: got %v expected %v", output, out)
	}

}

func TestChallenge6(t *testing.T) {
	in, err := ioutil.ReadFile("./inputs/6.txt")
	if err != nil {
		t.Fatalf("failed to read input; %v", err)
	}

	data, err := base64.StdEncoding.DecodeString(string(in))
	if err != nil {
		t.Fatalf("failed to decode from base64; %v", err)
	}

	keysize := decryption.FindKeySize(data, 40)
	transposed := decryption.PrepareInput(data, keysize)
	output := decryption.Solve(data, transposed, keysize)
	t.Logf(string(output))
}

func TestChallenge7(t *testing.T) {
	c, err := aes.NewCipher([]byte("YELLOW SUBMARINE"))
	if err != nil {
		t.Fatalf("failed to create aes cipher; %v", err)
	}

	in, err := ioutil.ReadFile("./inputs/7.txt")
	if err != nil {
		t.Fatalf("could not read input; %v", err)
	}

	in, err = base64.RawStdEncoding.DecodeString(string(in))
	if err != nil {
		t.Fatalf("failed to decode from base64; %v", err)
	}

	out, err := decryption.DecryptECD(in, c)
	if err != nil {
		t.Fatalf("failed to aes ecb decrypt; %v", err)
	}

	t.Log(string(out))
}

func TestChallenge8(t *testing.T) {
	in := getInput(t, "./inputs/8.txt")
	for i, bytes := range in {
		ecb, err := decryption.DetectECB(bytes, 16)
		if err != nil {
			t.Fatalf("failed to check if input is ECB encoded; %v", err)
		}
		if ecb {
			t.Logf("Line: %d is probably ECB encoded\n", i+1)
		}
	}
}

func getInput(t *testing.T, path string) [][]byte {
	var input [][]byte

	f, err := os.Open(path)
	if err != nil {
		t.Fatalf("failed to open file %s; %v", path, err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		bytes := toBytes(t, scanner.Text())
		input = append(input, bytes)
	}

	if err := scanner.Err(); err != nil {
		t.Fatalf("failed to scan input; %v", err)
	}

	return input
}

func toBytes(t *testing.T, in string) []byte {
	input, err := hex.DecodeString(in)
	if err != nil {
		t.Fatalf("failed to convert %s to bytes; %v", in, err)
	}
	return input
}
