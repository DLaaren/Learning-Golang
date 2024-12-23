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

type ThreadSafeCache struct {
	cache map[string]int
	mu    sync.Mutex
}

func InitThreadSafeCache() *ThreadSafeCache {
	cache := ThreadSafeCache{cache: make(map[string]int)}
	return &cache
}

func (c *ThreadSafeCache) Set(key string, value int) {
	c.mu.Lock()
	c.cache[key] = value
	c.mu.Unlock()
}

func (c *ThreadSafeCache) IsPresent(key string) bool {
	c.mu.Lock()
	_, ok := c.cache[key]
	defer c.mu.Unlock()
	return ok
}

var wg sync.WaitGroup

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cache *ThreadSafeCache) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	cache.Set(url, depth)

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		ok := cache.IsPresent(u)
		if !ok {
			wg.Add(1)
			go Crawl(u, depth-1, fetcher, cache)
		}
	}

	return
}

func main() {
	wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher, InitThreadSafeCache())
	wg.Wait()
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
