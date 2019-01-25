package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"os"
)

func fetch(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("http.Get: %s: %v\n", url, err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// exercise 1.9
	fmt.Printf("%v\n", resp.Status)

	// exercise 1.7
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		fmt.Printf("io.Copy: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	for _, url := range os.Args[1:] {
		// exercise 1.8
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}

		fetch(url)
	}
}
