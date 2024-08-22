package main

func ProcessCrawledLink(sitesChannel chan string, crawledLinksChannel chan string, pendingSitesChannel chan int) {

	foundUrls := make(map[string]bool)

	for c := range crawledLinksChannel {

		if len(foundUrls) >= 50 {
			close(sitesChannel)
			close(crawledLinksChannel)
			close(pendingSitesChannel)
			return
		}
		
		if !foundUrls[c] {
			foundUrls[c] = true
			pendingSitesChannel <- 1
			sitesChannel <- c
		}
	}
}