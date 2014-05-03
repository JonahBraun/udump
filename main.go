/*
*/

package main

import (
	"bytes"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	in := os.Stdin

	b := make([]byte, 1)

	for{
		r := make([]byte, 0, utf8.UTFMax)

		for{
			n, err := in.Read(b)

			if err != nil {
				if err.Error() == "EOF" {
					os.Exit(0)
				}
				panic(err)
			}
			if n<1 {
				panic("read less than 1 byte")
			}

			r = append(r, b[0])
			if utf8.Valid(r) {
				ru :=bytes.Runes(r)[0] 
				fmt.Printf("%U %c\n", ru, ru )
				break
			}
		}
	}
}
