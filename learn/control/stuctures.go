package main

import (
	"fmt"
)

func main() {

	atox := "The quick brown fox jumps over the lazy dog"
	vowels := 0
	consinent := 0

	for _, r := range atox {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			vowels += 1
		default:
			consinent += 1
		}
	}
	fmt.Printf("%d : %d \n", vowels, consinent)
}
