package main

import (
	"regexp"
	"strings"
)

// Functions in this script aims to increase the precision of BioHits by
// 1. removing "s" or "es" in Plurals, which transforms them into singulars
// 2.

// HandlePlural
func HandlePlural(wordFreq map[string]int) map[string]int {
	allWords := GetKeys(wordFreq)
	// check whether the word without es or s exists in the map
	// if true, remove this plural from the map
	for w := range wordFreq {
		if strings.HasSuffix(w, "s") {
			sChar := regexp.MustCompile(`s$`)
			esChar := regexp.MustCompile(`es$`)
			removeS := sChar.ReplaceAllString(w, "")
			removeES := esChar.ReplaceAllString(w, "")
			if StringInList(removeS, allWords) {
				wordFreq[removeS] += wordFreq[w]
				wordFreq[w] = 0
			} else if StringInList(removeES, allWords) {
				wordFreq[removeES] += wordFreq[w]
				wordFreq[w] = 0
			}
		}
	}

	// remove words with 0 frequency by first creating a new map
	// and add word with fequency that is not 0
	newMap := make(map[string]int, 0)
	for k, v := range wordFreq {
		if v != 0 {
			newMap[k] = v
		}
	}
	return newMap
}
