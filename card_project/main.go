package main

import "fmt"

func main() {

	var numbers [11]int

	for i := range numbers {
		numbers[i] = i
	}

	for _, n := range numbers {
		if n % 2 == 0 {
			fmt.Printf("%v is even\n", n)
		} else {
			fmt.Printf("%v is odd\n", n)
		}
	}

	//cards := newDeckFromFile("my_cards")
	//cards := newDeck()
	//cards.shuffle()
	//cards.print()
	//cards := newDeck()
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")
}
