package main

import (
	"fmt"
	"sync"
	"time"
)

// Fetcher returns the body of URL and a slice of URLs found on that page.
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

// SafeMap thread safe map
type SafeMap struct {
	urls map[string]int
	mux  sync.Mutex
}

func (m *SafeMap) visited(url string) bool {
	m.mux.Lock()
	defer m.mux.Unlock()
	_, exists := m.urls[url]
	if exists {
		return true
	}
	m.urls[url] = 1
	return false
}

// Crawl uses fetcher to recursively crawl pages starting with url, to a maximum of depth.
func (m *SafeMap) Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}

	if m.visited(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		go m.Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	m := SafeMap{urls: make(map[string]int)}
	go m.Crawl("https://golang.org/", 4, fetcher)
	time.Sleep(1000 * time.Millisecond)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
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
