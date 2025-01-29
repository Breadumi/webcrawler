package main

import (
	"fmt"
	URL "net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {

	cfg.concurrencyControl <- struct{}{}
	defer func() { <-cfg.concurrencyControl }()
	defer cfg.wg.Done()

	cfg.mu.Lock()
	pageLength := len(cfg.pages)
	cfg.mu.Unlock()
	if pageLength >= cfg.maxPages {
		return
	}

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
			urlparts, err := URL.Parse(url)
			if err != nil {
				fmt.Println(err)
			}
			if urlparts.Host != cfg.baseURL { // do not crawl links to external sites
				continue
			}
			cfg.wg.Add(1)
			go cfg.crawlPage(url)
		}
	}
}

func (cfg *config) addPageVisit(url string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if len(cfg.pages) >= cfg.maxPages {
		return false
	}
	if _, ok := cfg.pages[url]; !ok {
		isFirst = true
		cfg.pages[url] = 1
	} else {
		isFirst = false
		cfg.pages[url] += 1
	}
	return isFirst
}
