package main

import "fmt"

func main() {

	//var colors map[string]string

	var colors = make(map[string]string)

	//colors := map[string]string{
	//	"red": "#ff0000",
	//	"green": "#00ff00",
	//}
	colors["red"] = "#ff0000"
	colors["white"] = "#ffffff"
	fmt.Println(colors)
	//delete(colors, "red")
	fmt.Println(colors)
	printMap(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Printf("%v %v\n", color, hex)
	}
}
