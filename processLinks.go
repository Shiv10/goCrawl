package main

func ProcessCrawledLink(sitesChannel chan string, crawledLinksChannel chan string, pendingSitesChannel chan int) {

	foundUrls := make(map[string]bool)

	for c := range crawledLinksChannel {
		if !foundUrls[c] {
			foundUrls[c] = true
			pendingSitesChannel <- 1
			sitesChannel <- c
		}
	}
}