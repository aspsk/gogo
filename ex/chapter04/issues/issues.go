package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {

	params := "&per_page=100"

	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q + params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func prIssues(format string, items []*Issue) {
	if len(items) > 0 {
		fmt.Printf(format, len(items))
		for _, item := range items {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	monthAgo := now.AddDate(0, -1, 0)
	yearAgo := now.AddDate(-1, 0, 0)

	var month, year, older []*Issue
	for _, item := range result.Items {
		if !item.CreatedAt.Before(monthAgo) {
			month = append(month, item)
		} else if !item.CreatedAt.Before(yearAgo) {
			year = append(year, item)
		} else {
			older = append(older, item)
		}
	}

	fmt.Printf("Total: %d issues:\n", result.TotalCount)
	prIssues("%d issues less than a month old:\n", month)
	prIssues("%d issues less than a year old:\n", year)
	prIssues("%d issues more than a year old:\n", older)
}
