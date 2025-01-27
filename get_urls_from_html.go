package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {

	normURL, err := normalizeURL(rawBaseURL)
	if err != nil {
		return nil, err
	}

	var links []string

	htmlReader := strings.NewReader(htmlBody)
	doc, err := html.Parse(htmlReader)
	if err != nil {
		return nil, err
	}
	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					if _, err := normalizeURL(a.Val); err != nil {
						if len(a.Val) < 2 {
							log.Fatal("unexpected anchor element (length less than 2)")
						}
						links = append(links, fmt.Sprintf("https://%v%v", normURL, a.Val[1:]))
					} else {
						links = append(links, a.Val)
					}

					break
				}
			}
		}
	}
	return links, nil
}
