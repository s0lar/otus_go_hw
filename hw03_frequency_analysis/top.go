package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type WordCount struct {
	Word  string
	Count int
}

func Top10(input string) []string {
	words := strings.Fields(input)

	wordsCounted := []*WordCount{}
	ptrWords := make(map[string]*WordCount)

	for _, word := range words {
		if _, ok := ptrWords[word]; !ok {
			newWordCount := WordCount{word, 0}
			ptrWords[word] = &newWordCount
			wordsCounted = append(wordsCounted, &newWordCount)
		}
		ptrWords[word].Count++
	}

	sort.Slice(wordsCounted, func(i, j int) bool {
		return (wordsCounted[i].Count > wordsCounted[j].Count) ||
			((wordsCounted[i].Count == wordsCounted[j].Count) && wordsCounted[i].Word < wordsCounted[j].Word)
	})

	top10 := []string{}

	for i, word := range wordsCounted {
		if i >= 10 {
			break
		}
		top10 = append(top10, word.Word)
	}

	return top10
}
