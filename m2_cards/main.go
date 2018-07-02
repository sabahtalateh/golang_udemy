package main

import "fmt"

func main() {
	cards := [] string {"1", getZoz()}
	cards = append(cards, "Zog")
	fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func getZoz() string {
	return "Zoz"
}
