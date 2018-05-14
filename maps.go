package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wm := make(map[string]int)
	for _, w := range strings.Fields(s) {
		wm[w] += 1
	}
	return wm
}

func main() {
	wc.Test(WordCount)
}