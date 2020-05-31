package conversion

// CorpusEN according to https://www.sttmedia.com/characterfrequency-english
var CorpusEN = map[byte]float32{
	65:  8.34,
	66:  1.54,
	67:  2.73,
	68:  4.14,
	69:  12.60,
	70:  2.03,
	71:  1.92,
	72:  6.11,
	73:  6.71,
	74:  0.23,
	75:  0.87,
	76:  4.24,
	77:  2.53,
	78:  6.80,
	79:  7.70,
	80:  1.66,
	81:  0.09,
	82:  5.68,
	83:  6.11,
	84:  9.37,
	85:  2.85,
	86:  1.06,
	87:  2.34,
	88:  0.20,
	89:  2.04,
	90:  0.06,
	97:  8.34,
	98:  1.54,
	99:  2.73,
	100: 4.14,
	101: 12.60,
	102: 2.03,
	103: 1.92,
	104: 6.11,
	105: 6.71,
	106: 0.23,
	107: 0.87,
	108: 4.24,
	109: 2.53,
	110: 6.80,
	111: 7.70,
	112: 1.66,
	113: 0.09,
	114: 5.68,
	115: 6.11,
	116: 9.37,
	117: 2.85,
	118: 1.06,
	119: 2.34,
	120: 0.20,
	121: 2.04,
	122: 0.06,
	32:  10.00, // Adding heavy weight on space character
}

// FindSingleXorKey checks the score of a given text when xored with each
// ascii char from 0 to 127. It returns the maxScore of the text and the
// corresponding key.
func FindSingleXorKey(text []byte) (float32, byte) {
	var (
		char, key       byte
		score, maxScore float32
	)

	for char = 0; char < 128; char++ {
		for i := 0; i < len(text); i++ {
			v := text[i] ^ char
			score += CorpusEN[v]
		}
		if score > maxScore {
			maxScore = score
			key = char
		}
		score = 0
	}

	return maxScore, key
}
