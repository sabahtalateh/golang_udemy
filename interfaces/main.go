package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func main() {
	printGreeting(englishBot{})
	printGreeting(spanishBot{})
}

func printGreeting(bot bot) {
	fmt.Println(bot.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hello guy!";
}

func (sb spanishBot) getGreeting() string {
	return "Hola amigo!";
}
