package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            string
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		log.Fatal("no website provided")
	}
	if len(args) < 3 {
		log.Fatal("usage: URL maxConcurrency maxPages")
	}
	if len(args) > 3 {
		log.Fatal("too many arguments provided")
	}

	baseURL := args[0]
	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal("maxConcurrency must be an integer")
	}
	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		log.Fatal("maxPages must be an integer")
	}

	fmt.Printf("starting crawl of: %v\n", baseURL)

	pages := make(map[string]int)
	pages[args[0]] = 1

	cfg := config{
		pages:              pages,
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(cfg.baseURL)
	cfg.wg.Wait()

	cfg.printReport()

}

func (cfg *config) printReport() {
	fmt.Printf("=============================\n  REPORT for %s\n=============================\n", cfg.baseURL)
	sortedKeys := cfg.sortPages()

	for _, k := range sortedKeys {
		fmt.Printf("Found %v internal links to %s\n", cfg.pages[k], k)
	}

}

func (cfg *config) sortPages() []string {
	// return slice of keys sorted by their associated values

	type pair struct {
		key string
		val int
	}

	keyValues := make([]pair, 0, len(cfg.pages))

	for k, v := range cfg.pages {
		keyValues = append(keyValues, pair{key: k, val: v})
	}

	sort.Slice(keyValues, func(i, j int) bool {
		return keyValues[i].val > keyValues[j].val
	})

	keys := make([]string, 0, len(cfg.pages))
	for _, kv := range keyValues {
		keys = append(keys, kv.key)
	}

	return keys

}
