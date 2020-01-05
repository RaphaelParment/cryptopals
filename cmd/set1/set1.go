package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"

	c "raphael.parment/cryptopals/pkg/conversion"
)

func main() {
	// if err := challenge3(); err != nil {
	// 	fmt.Println("challenge 3 FAILED")
	// }

	// if err := challenge4(); err != nil {
	// 	fmt.Println("challenge 4 FAILED")
	// }
}

func challenge3() error {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	in, err := hex.DecodeString(input)
	if err != nil {
		return fmt.Errorf("failed to decode string '%s'", err.Error())
	}

	key, _ := c.FindKeyScore(in)
	for i := range in {
		fmt.Printf("%s", string(c.Xor(in[i], key)))
	}
	fmt.Println()

	return nil
}

func challenge4() error {
	var (
		input    []string
		maxScore float32
		msgKey   byte
		msgIn    []byte
	)

	// Read the input
	f, err := os.Open("./inputs/4.txt")
	if err != nil {
		return fmt.Errorf("failed to open file '%s'", err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to scan input '%s'", err.Error())
	}

	for i, line := range input {
		in, err := hex.DecodeString(line)
		if err != nil {
			return fmt.Errorf("failed to decode line %d '%s'", i, err.Error())
		}
		key, score := c.FindKeyScore(in)
		if score > maxScore {
			maxScore = score
			msgKey = key
			msgIn = in
		}
	}

	for i := range msgIn {
		fmt.Printf("%s", string(c.Xor(msgIn[i], msgKey)))
	}

	return nil
}
