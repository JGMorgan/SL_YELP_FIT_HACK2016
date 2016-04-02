package chatfunctions

import (
	"fmt"
	"math/rand"
	"time"
)

var options []string

func yelp() []byte {
	return ([]byte("functionality in development"))
}

func swag() []byte {
	return ([]byte("SWAG\nSWAG\nSWAG\nSWAG\nSWAG\nSWAG\nSWAG\nSWAG\nSWAG\nSWAG\nSWAG\nSWAG\nSWAG\nSWAG\nSWAG\nSWAG\n"))
}

func flipcoin() []byte {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	x := r.Int() % 2
	if x == 0 {
		return ([]byte("Heads"))
	} else {
		return ([]byte("Tails"))
	}
}

func pickforme() []byte {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	x := r.Int() % len(options)
	return ([]byte(options[x]))
}
