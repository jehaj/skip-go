package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type player struct {
	stack  stack
	stacks [4]stack
	hand   [5]int
}

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

	const number_of_player_cards = 20
	const number_of_players = 1
	const number_of_cards_on_hand = 5
	all_cards_pointer := len(all_cards) - 1
	players := make([]player, number_of_players)
	for i := 0; i < number_of_players; i++ {
		for j := 0; j < number_of_player_cards; j++ {
			players[i].stack.push(get_card(all_cards[:], &all_cards_pointer))
		}
	}

	var out_stacks [4]stack
	player := players[0]
	done := false
	for !done {
		fmt.Printf("You have %d cards in your stack with %d being on top.\n", len(player.stack), player.stack.peek())
		cards_to_draw := 0
		for i, v := range player.hand {
			if v == 0 {
				player.hand[i] = get_card(all_cards[:], &all_cards_pointer)
				cards_to_draw += 1
			}
		}
		cards_on_hand := make([]string, number_of_cards_on_hand)
		for i, v := range player.hand {
			cards_on_hand[i] = strconv.Itoa(v)
		}
		fmt.Printf("You draw %d cards and on your hand is %s.\n", cards_to_draw, strings.Join(cards_on_hand, " "))
		fmt.Println("Which card do you want to use (0-5)?")
		var card_to_use_input string
		fmt.Scanln(&card_to_use_input)
		card_to_use_index, _ := strconv.ParseInt(card_to_use_input, 10, 64)
		var card_to_use int
		if card_to_use_index == 0 {
			card_to_use = player.stack.pop()
		} else {
			card_to_use = player.hand[card_to_use_index-1]
			player.hand[card_to_use_index-1] = 0
		}
		fmt.Printf("Where should %d be placed (0-7)?\n", card_to_use)
		var position_for_card_string string
		fmt.Scanln(&position_for_card_string)
		position_for_card, _ := strconv.ParseInt(position_for_card_string, 10, 64)
		if position_for_card < 4 {
			out_stacks[position_for_card%4].push(card_to_use)
		} else {
			player.stacks[position_for_card%4].push(card_to_use)
			done = true
			fmt.Println("You have placed a card in your own stack, so your turn is now over!")
		}
	}
	fmt.Println(all_cards)
}

func get_card(card_stack []int, card_stack_pointer *int) int {
	r := card_stack[*card_stack_pointer]
	card_stack[*card_stack_pointer] = 0
	*card_stack_pointer -= 1
	return r
}
