package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	resultsChan := make(chan result)
	results := make(map[string]bool)

	for _, url := range urls {
		resultsChan <- result{url, wc(url)}
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultsChan
		results[r.string] = r.bool

	}

	return results
}
