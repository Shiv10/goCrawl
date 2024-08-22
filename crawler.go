package main

import (
	"fmt"
	"sync"
)

func Crawl(wg *sync.WaitGroup, sitesChannel chan string, crawledLinksChannel chan string, pendingSitesChannel chan int) {
	defer wg.Done()

	for site := range sitesChannel {
		fmt.Printf("Found website: \n%s\n", site)
		extractLink(site, crawledLinksChannel)
		pendingSitesChannel <- -1
	}
}