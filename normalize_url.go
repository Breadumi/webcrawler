package main

import (
	"errors"
	"fmt"
	URL "net/url"
	"strings"
)

func normalizeURL(url string) (string, error) {
	parsedURL, err := URL.Parse(url)
	if err != nil {
		return "", err
	}

	if strings.ToLower(url[0:8]) != "https://" && strings.ToLower(url[0:7]) != "http://" {
		return "", errors.New("malformed URL")
	}

	fmt.Println(url)
	host := parsedURL.Host
	path := parsedURL.Path
	fmt.Println(host)
	fmt.Println(path)
	var normURL string

	hostValid := true
	pathValid := true

	if len(path) > 1 && strings.HasSuffix(path, "//") {
		pathValid = false
	}
	if len(host) > 0 && strings.HasPrefix(host, "/") {
		hostValid = false
	}
	if host == "" {
		hostValid = false
	}

	if !pathValid || !hostValid {
		fmt.Println(pathValid, hostValid)
		fmt.Println("")
		return "", errors.New("malformed URL")
	}

	if path == "" {
		normURL = fmt.Sprintf("%v/", host)
	} else {
		normURL = fmt.Sprintf("%v%v", host, path)
	}
	fmt.Println(normURL)
	fmt.Println("")
	return normURL, nil

}
