package main

import (
	"fmt"
	"math/rand"
)

type stack []int

func (s *stack) push(e int) {
	*s = append(*s, e)
}

func (s *stack) peek() int {
	return (*s)[len(*s)-1]
}

func (s *stack) pop() int {
	e := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return e
}

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

	s1 := make(stack, 0)
	s1.push(all_cards[10])
	fmt.Println(s1)
	fmt.Println(s1.pop())
	fmt.Println(s1)

	fmt.Println("Hello World")
	fmt.Println(all_cards)

}
