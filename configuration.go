package main

import (
	"math/rand"
	"time"
)

type Configuration struct {
	Length   int
	Dashed   bool
	Alphabet []rune
}

var config *Configuration
var defaultAlphabet = []rune("QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm1234567890'^()/&%$£\"!")

func shuffle(alphabet *[]rune) {
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(*alphabet), func(i, j int) {
		(*alphabet)[i], (*alphabet)[j] = (*alphabet)[j], (*alphabet)[i]
	})
}

func newConfiguration(length int, dashed bool, alphabet *[]rune) *Configuration {
	return &Configuration{
		Length:   length,
		Dashed:   dashed,
		Alphabet: *alphabet,
	}
}
