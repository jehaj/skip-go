package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var all_cards [162]int
	for i := 0; i < 12*12; i++ {
		all_cards[i] = i%12 + 1
	}
	for i := 12 * 12; i < 162; i++ {
		all_cards[i] = -1
	}

	rand.Shuffle(len(all_cards), func(i, j int) {
		all_cards[i], all_cards[j] = all_cards[j], all_cards[i]
	})

	fmt.Print("Hello World")
	fmt.Print("{}", all_cards)
}
