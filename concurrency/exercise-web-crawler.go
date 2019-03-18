package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type URLMux struct {
	visited map[string]bool
	mux     sync.Mutex
}

func CrawlParalell(url string, depth int, fetcher Fetcher, ret chan string, visited *URLMux) {
	defer close(ret)

	if depth <= 0 {
		return
	}
	visited.mux.Lock()
	if visited.visited[url] {
		visited.mux.Unlock()
		return
	}

	visited.visited[url] = true
	visited.mux.Unlock()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	ret <- fmt.Sprintf("found: %s %q\n", url, body)

	result := make([]chan string, len(urls))
	for i, u := range urls {
		result[i] = make(chan string)
		go CrawlParalell(u, depth-1, fetcher, result[i], visited)
	}

	for i := range result {
		for s := range result[i] {
			ret <- s
		}
	}

	return

}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	result := make(chan string)
	visited := URLMux{visited: make(map[string]bool)}
	go CrawlParalell(url, 4, fetcher, result, &visited)

	for s := range result {
		fmt.Println(s)
	}

}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	fmt.Println("fetchin", url)
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
