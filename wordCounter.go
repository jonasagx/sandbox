package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordsCounter := make(map[string]int)
	
	for _,word := range(strings.Fields(s)) {
		counted, isIn := wordsCounter[word]
		if isIn {
			wordsCounter[word] = counted + 1
		} else {
			wordsCounter[word] = 1
		}
	}
	return wordsCounter
}

func main() {
	wc.Test(WordCount)
}
