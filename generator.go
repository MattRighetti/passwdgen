package passwdgen

import (
	"math/rand"
	"time"
)

const (
	// LENGTH

	// L8 length 8
	L8 = uint8(8)
	// L10 length 10
	L10 = uint8(10)
	// L15 length 15
	L15 = uint8(15)
	// L16 length 16
	L16 = uint8(16)
	// L20 length 20
	L20 = uint8(20)
	// L32 length 32
	L32 = uint8(32)
	// L64 length 64
	L64 = uint8(64)

	// KINDS

	DashSeparated = "ds"
	Normal        = "n"
)

// Generate executes a function that will return a dash-separated password of length and kind
func Generate(length uint8) []byte {
	seq := []byte{}
	switch length {
	case L8, L16, L32, L64:
		for i := 0; i < int(length)/4; i++ {
			seq = append(seq, generateRandomSequence(4)...)
			seq = append(seq, '-')
		}
	case L10, L15, L20:
		for i := 0; i < int(length)/5; i++ {
			seq = append(seq, generateRandomSequence(5)...)
			seq = append(seq, '-')
		}
	}

	return seq[:len(seq)-1]
}

func generateRandomSequence(length uint8) []byte {
	seq := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(seq); i++ {
		seq[i] = pickRandom(rand.Int31n(3), rand.Int31())
	}

	return seq
}

func pickRandom(i int32, j int32) byte {
	numbers := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	letters := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	upperLetters := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

	matrix := [][]byte{numbers, letters, upperLetters}

	var rByte byte
	if i == 0 {
		rByte = matrix[0][j%10]
	} else {
		rByte = matrix[i][j%26]
	}

	return rByte
}
