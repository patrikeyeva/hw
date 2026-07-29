package hw03frequencyanalysis

import (
	"slices"
	"strings"
	"unicode"
)

const (
	limit = 10
	dash  = "-"
)

type Word struct {
	word  string
	count uint
}

func cmp(a, b Word) int {
	if a.count == b.count {
		if a.word < b.word {
			return -1
		} else if a.word > b.word {
			return 1
		}
	}
	return int(b.count - a.count)
}

func cleanWord(word string) string {
	runes := []rune(strings.ToLower(word))
	first := 0

	for len(runes) > 0 && unicode.IsPunct(runes[first]) {
		runes = runes[first+1:]
	}
	for len(runes) > 0 && unicode.IsPunct(runes[len(runes)-1]) {
		runes = runes[:len(runes)-1]
	}

	return string(runes)
}

func Top10(s string) []string {
	str := strings.Fields(s)

	counter := make(map[string]uint)
	for _, word := range str {
		word := cleanWord(word)
		if word != dash && len(word) > 0 {
			counter[word]++
		}
	}

	words := make([]Word, len(counter))
	i := 0
	for word, count := range counter {
		words[i].word, words[i].count = word, count
		i++
	}

	slices.SortFunc(words, cmp)

	first := words[:min(limit, len(words))]
	res := make([]string, min(limit, len(words)))
	for i, word := range first {
		res[i] = word.word
	}

	return res
}
