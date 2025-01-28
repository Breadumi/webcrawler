package main

import (
	"fmt"
)

func (cfg *config) crawlPage(rawCurrentURL string) {

	cfg.concurrencyControl <- struct{}{}
	defer func() { <-cfg.concurrencyControl }()
	defer cfg.wg.Done()

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Scanning: %v\n", rawCurrentURL)

	newURLs, err := getURLsFromHTML(html, cfg.baseURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Found %v:\n", newURLs)

	for _, url := range newURLs {
		if cfg.addPageVisit(url) {
			cfg.wg.Add(1)
			go cfg.crawlPage(url)
		}
	}
}

func (cfg *config) addPageVisit(url string) (isFirst bool) {
	cfg.mu.Lock()
	if _, ok := cfg.pages[url]; !ok {
		isFirst = true
		cfg.pages[url] = 1
	} else {
		isFirst = false
		cfg.pages[url] += 1
	}
	cfg.mu.Unlock()
	return isFirst
}
