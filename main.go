// -----------------------------------------------------------------------------
// Encryption Demo Project                                     gocrypt/[main.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package main

import (
	"fmt"

	"github.com/balacode/zr"
	whirl "github.com/balacode/zr-whirl"
)

var (
	PL = fmt.Println
)

// main _ _
func main() {
	var (
		key1 = []byte("fryling and grimbling nambs")
		key2 = []byte("Jabbywock!")
		t    zr.Timer
	)
	t.Start("Generate 1 million hashes")
	{
		multiHash(key1, key2, 1e6)
	}
	t.StopLast()
	t.Print()
}

// multiHash _ _
func multiHash(data, key []byte, iterations int) []byte {
	var (
		salt = key
		hash = whirl.HashOfBytes(data, salt)
	)
	for i := 0; i < iterations; i++ {
		hash = whirl.HashOfBytes(hash, salt)
	}
	return hash
}

/*
SN  Seconds:
#01 5.58937
#02 5.64750
#03 5.69256
#04 5.69947
#05 5.72258
#06 5.77542
#07 5.77642
#08 5.79238
#09 5.81740
#10 5.84438
*/

// end
