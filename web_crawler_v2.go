package main

import (
	"fmt"
	"errors"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type FetchR struct {
	url string
	urls []string
	body string
	err error
	depth int
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	urlAccessed = make(map[string]bool)
	task := 0
	fetch_chan := make(chan *FetchR)
	fetch := func(url string, ch chan *FetchR, depth int) {
		var body string
		var urls []string
		var err error
		if depth >= 0 {
			body, urls, err = fetcher.Fetch(url)
		} else {
			err = errors.New("Max depth reached")
		}
		ch <- &FetchR{url, urls, body, err, depth}
	}

	go fetch(url, fetch_chan, depth)
	urlAccessed[url] = true
	task++
	
	for {
		re := <- fetch_chan
		task--
		if re.err != nil {
			fmt.Println(re.err)
		} else {
			fmt.Printf("found: %s %q\n", re.url, re.body)
			for _, u := range re.urls {
				if !urlAccessed[u] {
					go fetch(u, fetch_chan, re.depth-1)
					urlAccessed[u] = true
					task++
				}
			}
		}
		if task == 0 {
			close(fetch_chan)
			break
		}
	}
}

func main() {
	Crawl("https://golang.org/", 1, fetcher)
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

var urlAccessed map[string]bool

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