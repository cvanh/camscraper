package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	getCatagory()
}

// get list of products by catagory number(125 for nut and cams on klimwinkel)
func getCatagory() {
	client := &http.Client{}
	var data = strings.NewReader(`{
	"alternative_options_photos": true,
	"category_id": 127,
	"brands_filter": [],
	"stock_filter": [],
	"specs_filter": [],
	"sort": "default"
	}`)

	req, err := http.NewRequest("POST", "https://www.klimwinkel.nl/api/queryProducts", data)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	// TODO make shure these wont expire
	req.Header.Set("Cookie", "fzshoprest_session=eyJpdiI6Ik1sR0tnQlBOU0wwL2ZBdHU5bm1NZXc9PSIsInZhbHVlIjoiZm1XZE5MeGE5Q1hTRjdnQTFzSU1IQ0hDTDlEUUtZUzg5WTh6VnhUSE9xS3RKdDVUK1I0YjltbU1rOE91TmcvOUJpVnNydDd3SVY2TnlHUmgxdnhUdzd5OG1saDRiVUJwTWcreTQ3ZHFtQU9hZHM3a08xM1BPdjRBNm8zUkpURk8iLCJtYWMiOiIxZjkyYzgzNTM5YTk4MzRlZWQ1MjYwYjYwODU2OTAyZDQ3NmIyZmRkYmVhNTE1MjQ3ZTY0YjdhZjlmNjYzMTZhIiwidGFnIjoiIn0%253D; i18n_redirected=nl; auth.strategy=local")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", bodyText)
}
