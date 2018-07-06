package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards.print()
	//cards := newDeck()
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")
}
