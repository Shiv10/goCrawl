package main

func MonitorChannels(sitesChannel chan string, crawledLinksChannel chan string, pendingSitesChannel chan int) {
	count := 0

	for i := range pendingSitesChannel {
		count += i

		if count == 0 {
			close(sitesChannel)
			close(crawledLinksChannel)
			close(pendingSitesChannel)
		}
	}
}