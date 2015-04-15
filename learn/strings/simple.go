package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "asSASA ddd Dsjkdsj dk"
	length := len([]byte(str))
	bytes := utf8.RuneCount([]byte(str))

	fmt.Printf("String %s\nLength: %d, Runes: %d", str, length, bytes)

}
