package concurrency

// WebsiteChecker type
type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

// CheckWebsites is a function that check url in urls should be accessed
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}
	return results
}
