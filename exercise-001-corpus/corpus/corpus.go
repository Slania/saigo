package corpus

import (
	"regexp"
	"sort"
	"strings"
)

type WordCount struct {
	Word  string
	Count int
}

type Corpus []WordCount

func (corpus Corpus) Swap(i, j int)      { corpus[i], corpus[j] = corpus[j], corpus[i] }
func (corpus Corpus) Len() int           { return len(corpus) }
func (corpus Corpus) Less(i, j int) bool { return corpus[i].Count < corpus[j].Count }

func Analyze(data string) Corpus {
	dictionary := make(map[string]int)

	for _, element := range regexp.MustCompile(`\s+`).Split(string(data), -1) {
		word := regexp.MustCompile(`^[^\w]*|[^\w]*$`).ReplaceAllString(strings.ToLower(element), "")
		if _, ok := dictionary[strings.ToLower(word)]; ok {
			dictionary[strings.ToLower(word)]++
		} else {
			dictionary[strings.ToLower(word)] = 1
		}
	}

	corpus := make(Corpus, 0)
	for word, count := range dictionary {
		corpus = append(corpus, WordCount{word, count})
	}
	sort.Sort(sort.Reverse(corpus))

	return corpus
}
