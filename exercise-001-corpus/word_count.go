package main

import (
	"fmt"
	"github.com/enova/saigo/exercise-001-corpus/corpus"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s [filename]\n", os.Args[0])
		os.Exit(1)
	}

	fileName := os.Args[1]

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}

	analysis := corpus.Analyze(string(data))
	for _, wordCount := range analysis {
		fmt.Printf("%s %d\n", wordCount.Word, wordCount.Count)
	}
}
