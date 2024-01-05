package main

import (
	"fmt"
	"sort"
	"strings"
)

// WordInfo struct represents information about a word,
// including the word itself, the count of 'a' characters,
// and the length of the word.
type WordInfo struct {
	Word    string
	ACount  int
	Length  int
}

// customSort function sorts a list of words based on the number of 'a' characters
// in each word (in decreasing order). If two words have the same number of 'a's,
// they are then sorted by their lengths in decreasing order.
func customSort(words []string) []WordInfo {
	var wordInfos []WordInfo

	// Sort the words based on custom criteria
	sort.Slice(words, func(i, j int) bool {
		aCountI := strings.Count(words[i], "a")
		aCountJ := strings.Count(words[j], "a")

		if aCountI == aCountJ {
			// If the 'a' counts are equal, sort by length
			return len(words[i]) > len(words[j])
		}

		// Sort by the number of 'a' characters
		return aCountI > aCountJ
	})
	
	// Populate WordInfo struct for each word
	for _, word := range words {
		wordInfo := WordInfo{
			Word:    word,
			ACount:  strings.Count(word, "a"),
			Length:  len(word),
		}
		wordInfos = append(wordInfos, wordInfo)
	}

	return wordInfos
}

func test() {
	wordList := []string{"ana", "hello", "apple", "avocado", "watermelon"}
	sortedWords := customSort(wordList)

	fmt.Println("Sorted Words and Information:")
	for _, info := range sortedWords {
		fmt.Printf("Word: %-20s 'a' Count: %2d Length: %-2d\n", info.Word, info.ACount, info.Length)
	}
}

func main(){
    test()
}
