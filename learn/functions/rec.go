package main

import (
	"fmt"
)

func rec(i int) {
	if i == 10 {
		return
	}
	rec(i + 1)
	fmt.Printf("%d ", i)
}

func main() {
	rec(0)
}
