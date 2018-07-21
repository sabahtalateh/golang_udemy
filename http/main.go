package main

import (
	"net/http"
	"fmt"
	"os"
	"io"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//bs := make([]byte, 1024 * 1024)
	//a, _ := resp.Body.Read(bs)
	//fmt.Println(string(bs))
	//fmt.Println(a)
	//io.Copy(os.Stdout, resp.Body)
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Printf("Wrote %v bytes.\n", len(bs))
	return len(bs), nil
}
