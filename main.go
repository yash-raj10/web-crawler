package main

import (
	"crypto/tls"
	"fmt"
	"github.com/steelx/extractlinks"
	"net/http"
	"net/url"
	"os"
)

var (
	urlQueue   = make(chan string)
	hasCrawled = make(map[string]bool)

	config = &tls.Config{
		InsecureSkipVerify: true,
	}
	transport = &http.Transport{
		TLSClientConfig: config,
	}
	netClient = &http.Client{
		Transport: transport,
	}
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("URL is missing, e.g. https://github.com/yash-raj10")
		os.Exit(1)
	}
	baseUrl := args[0]

	crawlUrl(baseUrl)

	for href := range urlQueue {
		if !hasCrawled[href] && isSameDomain(baseUrl, href) {
			crawlUrl(href)
		}
	}
}

func crawlUrl(baseUrl string) {
	hasCrawled[baseUrl] = true
	fmt.Println("Crawling -> ", baseUrl)
	resp, err := netClient.Get(baseUrl)
	checkError(err)
	defer resp.Body.Close()

	//getting the links
	AnkerLinks, err := extractlinks.All(resp.Body)
	checkError(err)

	for _, link := range AnkerLinks {
		if link.Href == "" {
			continue
		}

		fixedURL := fixedURLS(baseUrl, link.Href)

		go func(url string) {
			urlQueue <- url
		}(fixedURL)

	}

}

func fixedURLS(baseUrl, link string) string {
	base, err := url.Parse(baseUrl)
	if err != nil {
		return ""
	}
	linkUrl, err := url.Parse(link)
	if err != nil {
		return ""
	}

	fixedURL := base.ResolveReference(linkUrl)
	return fixedURL.String()
}

func isSameDomain(baseUrl, link string) bool {
	base, err := url.Parse(baseUrl)
	if err != nil {
		return false
	}
	linkUrl, err := url.Parse(link)
	if err != nil {
		return false
	}

	if linkUrl.Host != base.Host {
		return false
	}

	return true
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
