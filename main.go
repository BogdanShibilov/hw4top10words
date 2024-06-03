package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

func GetTopTenWords(text string) []string {
	if len(text) == 0 {
		return []string{}
	}

	words := strings.Split(text, " ")
	wordsFrequency := getWordsFrequency(words)
	wordsSet := setFrom(words)
	sortDescByFrequencyWithSortSlice(wordsSet, wordsFrequency)

	if len(wordsSet) <= 10 {
		return wordsSet
	} else {
		return wordsSet[0:10]
	}
}

func getWordsFrequency(words []string) map[string]int {
	var wordsFrequency = make(map[string]int)
	for _, word := range words {
		wordsFrequency[strings.ToLower(word)] = wordsFrequency[strings.ToLower(word)] + 1
	}

	return wordsFrequency
}

func setFrom(slice []string) []string {
	var set []string

	for _, element := range slice {
		if !slices.Contains(set, strings.ToLower(element)) {
			set = append(set, strings.ToLower(element))
		}
	}

	return set
}

//func sortDescByFrequency(wordsSet []string, wordsFrequency map[string]int) {
//	slices.SortFunc(wordsSet, func(left, right string) int {
//		return cmp.Compare(wordsFrequency[right], wordsFrequency[left])
//	})
//}

func sortDescByFrequencyWithSortSlice(wordsSet []string, wordsFrequency map[string]int) {
	sort.Slice(wordsSet, func(i, j int) bool {
		leftWord := wordsSet[i]
		rightWord := wordsSet[j]
		return wordsFrequency[leftWord] > wordsFrequency[rightWord]
	})
}

func main() {
	fmt.Println(GetTopTenWords("q w e r t y y y t t Y e"))
}
