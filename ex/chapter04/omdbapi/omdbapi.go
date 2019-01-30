package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type searchResult struct {
	Title string
	Year string
	Genre string
}

func search(apikey, title string) (*searchResult, error) {

	resp, err := http.Get(fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&t=%s", apikey, title))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("query failed: %s", resp.Status)
	}

	var result searchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func main() {

	apikey := flag.String("key", "", "API key")
	title := flag.String("title", "", "Title")
	flag.Parse()

	result, err := search(*apikey, *title)
	if err != nil {
		log.Fatalf("search: %v", err)
	}
	fmt.Println("Title:", result.Title)
	fmt.Println("Year: ", result.Year)
	fmt.Println("Genre:", result.Genre)
}
