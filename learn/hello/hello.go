package main

import (
	"fmt"
)

func main() {
	list := []string{"a", "b", "c", "d", "e"}
	for k, v := range list {
		fmt.Printf("%n  %s\n", k, v)
	}
}
