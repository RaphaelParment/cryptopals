package conversion

var idxToCharBase64Table = map[uint8]uint8{
	0:  65,
	1:  66,
	2:  67,
	3:  68,
	4:  69,
	5:  70,
	6:  71,
	7:  72,
	8:  73,
	9:  74,
	10: 75,
	11: 76,
	12: 77,
	13: 78,
	14: 79,
	15: 80,
	16: 81,
	17: 82,
	18: 83,
	19: 84,
	20: 85,
	21: 86,
	22: 87,
	23: 88,
	24: 89,
	25: 90,
	26: 97,
	27: 98,
	28: 99,
	29: 100,
	30: 101,
	31: 102,
	32: 103,
	33: 104,
	34: 105,
	35: 106,
	36: 107,
	37: 108,
	38: 109,
	39: 110,
	40: 111,
	41: 112,
	42: 113,
	43: 114,
	44: 115,
	45: 116,
	46: 117,
	47: 118,
	48: 119,
	49: 120,
	50: 121,
	51: 122,
	52: 48,
	53: 49,
	54: 50,
	55: 51,
	56: 52,
	57: 53,
	58: 54,
	59: 55,
	60: 56,
	61: 57,
	62: 43,
	63: 47,
}

func HexToBase64(input []byte) []byte {
	totalBits := make([]uint8, len(input)*8)
	for i, b := range input {
		bits := UByteToBits(b)
		for j := 0; j < 8; j++ {
			totalBits[j+(i*8)] = bits[j]
		}
	}

	var output []uint8
	var char uint8
	var used uint8
	for i := 0; i < len(totalBits); i++ {
		if totalBits[i] == 1 {
			char += 1 << (5 - i%6)

		}
		used++
		if used == 6 {
			output = append(output, idxToCharBase64Table[char])
			char = 0
			used = 0
		}
	}

	return output
}