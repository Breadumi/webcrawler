package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            string
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func main() {

	maxConcurrency := 100

	args := os.Args[1:]

	if len(args) < 1 {
		log.Fatal("no website provided")
	}
	if len(args) > 1 {
		log.Fatal("too many arguments provided")
	}

	fmt.Printf("starting crawl of: %v\n", args[0])

	pages := make(map[string]int)
	pages[args[0]] = 1

	cfg := config{
		pages:              pages,
		baseURL:            args[0],
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(cfg.baseURL)
	cfg.wg.Wait()

	for key, val := range pages {
		fmt.Printf("Page: %s\n\tCount: %v\n", key, val)
	}

}
