package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	sliceWord := strings.Fields(s)
	mapaCount := make(map[string]int)
	for _, word := range sliceWord {
		valor, ok := mapaCount[word]
		if ok {
			mapaCount[word] = valor + 1
		} else {
			mapaCount[word] = 1
		}
	}
	return mapaCount
}

func main() {
	wc.Test(WordCount)
}
