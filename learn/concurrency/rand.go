package main

import (
	"fmt"
	"math/rand"
)

func makeRand(c chan int) {
	for {
		c <- rand.Intn(100)
	}
}

func main() {
	randoms := make(chan int)

	go makeRand(randoms)

	for n := range randoms {
		fmt.Printf("%d ", n)
	}

}
