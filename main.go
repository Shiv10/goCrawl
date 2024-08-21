package main

import (
	"fmt"
	"regexp"
	"strings"

	// "io/ioutil"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	siteToCrawl := "https://theuselessweb.com/"

	response, err := http.Get(siteToCrawl)
	
	if err != nil {
		fmt.Printf("Error while readin site %s", err)
		return
	}

	defer response.Body.Close()

	// bodyBytes, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	fmt.Printf("Error while readin bytes %s", err)
	// 	return
	// }
	
	tokenizer := html.NewTokenizer(response.Body)
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			return
		}

		token := tokenizer.Token()
		re := regexp.MustCompile(`href=".+"`)
		if tokenType == html.StartTagToken && token.Data == "a" {
			link := string(re.Find([]byte(token.String())))[6:]
			index := strings.Index(link, "\"")
			link = link[0:index]
			if !strings.HasPrefix(link, "http") {
				continue
			}
			fmt.Println(link)
		}
	}
}