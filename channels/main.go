package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://ya.ru",
		"http://amazon.com",
		"http://twitter.com",
		"http://microsoft.com",
		"http://google.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkURL(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(time.Second)
			checkURL(link, c)
		}(l)
	}
}

func checkURL(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(url, "is up.")
	c <- url
}
