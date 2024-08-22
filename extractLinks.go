package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func extractLink(website string, crawledLinksChannel chan string) {

	response, err := http.Get(website)

	if err != nil {
		fmt.Printf("Error while readin site %s", err)
		return
	}

	defer response.Body.Close()
	tokenizer := html.NewTokenizer(response.Body)

	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			return
		}

		token := tokenizer.Token()
		re := regexp.MustCompile(`href=".+"`)
		if tokenType == html.StartTagToken && token.Data == "a" {
			link := string(re.Find([]byte(token.String())))
			if len(link) < 7 {
				continue
			}
			link = link[6:]
			index := strings.Index(link, "\"")
			link = link[0:index]
			if !strings.HasPrefix(link, "http") {
				continue
			}
			go func() {
				crawledLinksChannel <- link
			}()
		}
	}
}