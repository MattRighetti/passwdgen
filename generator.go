package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	// LENGTH

	// L8 length 8
	L8 = int(8)
	// L10 length 10
	L10 = int(10)
	// L15 length 15
	L15 = int(15)
	// L16 length 16
	L16 = int(16)
	// L20 length 20
	L20 = int(20)
	// L32 length 32
	L32 = int(32)
	// L64 length 64
	L64 = int(64)
)

// Generate executes a function that will return a dash-separated password of length and kind
func Generate() (seq []rune, err error) {
	if config.Dashed {
		seq, err = generateDashedPassword(config.Length)
	} else {
		seq = generatePassword(config.Length)
	}

	return
}

func generateDashedPassword(length int) (seq []rune, err error) {
	switch length {
	case L8, L16, L32, L64:
		for i := 0; i < length/4; i++ {
			seq = append(seq, generateRandomSequence(4)...)
			seq = append(seq, '-')
		}
	case L10, L15, L20:
		for i := 0; i < length/5; i++ {
			seq = append(seq, generateRandomSequence(5)...)
			seq = append(seq, '-')
		}
	default:
		err = fmt.Errorf("Length not supported")
		return
	}

	seq = seq[:len(seq)-1]

	return
}

func generatePassword(length int) []rune {
	return generateRandomSequence(length)
}

func generateRandomSequence(length int) []rune {
	seq := make([]rune, length)
	rand.Seed(time.Now().UnixNano())
	var n int
	for i := 0; i < length; i++ {
		n = rand.Intn(len(config.Alphabet))
		seq[i] = config.Alphabet[n]
	}
	return seq
}

func main() {
	switch len(os.Args) {
	case 2:
		length, _ := strconv.Atoi(os.Args[1])
		config = newConfiguration(int(length), true, &defaultAlphabet)
	case 3:
		length, _ := strconv.Atoi(os.Args[1])
		dashed, _ := strconv.ParseBool(os.Args[2])
		config = newConfiguration(int(length), dashed, &defaultAlphabet)
	case 4:
		length, _ := strconv.Atoi(os.Args[1])
		dashed, _ := strconv.ParseBool(os.Args[2])
		alphabet := []rune(os.Args[3])
		config = newConfiguration(int(length), dashed, &alphabet)
	default:
		config = newConfiguration(20, true, &defaultAlphabet)
	}

	pass, _ := Generate()

	fmt.Printf("%s\n", string(pass))
}
