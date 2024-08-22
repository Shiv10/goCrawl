package main

import "sync"

func main() {
	siteToCrawl := "https://theuselessweb.com/"

	sitesChannel := make(chan string)
	crawledLinksChannel := make(chan string)
	pendingSitesChannel := make(chan int)

	var wg sync.WaitGroup

	go func() {
		crawledLinksChannel <- siteToCrawl
	}()

	go ProcessCrawledLink(sitesChannel, crawledLinksChannel, pendingSitesChannel)
	go MonitorChannels(sitesChannel, crawledLinksChannel, pendingSitesChannel)

	numOfThreads := 50

	for i := 0; i < numOfThreads; i++ {
		wg.Add(1)
		go Crawl(&wg, sitesChannel, crawledLinksChannel, pendingSitesChannel)
	}

	wg.Wait()
}