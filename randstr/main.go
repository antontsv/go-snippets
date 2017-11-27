package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("Random base16: %s\n", RandBase16(64))
}

var base16 = []rune("abcdef0123456789")

// RandBase16 gives random string on len n using only base16 characters
func RandBase16(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = base16[rand.Intn(len(base16))]
	}
	return string(b)
}
