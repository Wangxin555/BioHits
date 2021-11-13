package main

import (
	"regexp"
	"strings"
)

// Function in this script aims to increase the precision of BioHits by
// 1. removing "s" or "es" in Plurals, which transforms them into singulars noun
// 2. removing "ly" in adv, which transform them into a normal sense verb

// WordTransform
func WordTransform(wordFreq map[string]int) map[string]int {
	allWords := GetKeys(wordFreq)
	for w := range wordFreq {
		// check whether the word without es or s exists in the map
		// if true, remove this plural from the map
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

		// check whether the word without ly exists in the map
		// if true, remove this adverb from the map
		if strings.HasSuffix(w, "ly") {
			lyChar := regexp.MustCompile(`ly$`)
			removeLy := lyChar.ReplaceAllString(w, "")
			if StringInList(removeLy, allWords) {
				wordFreq[removeLy] += wordFreq[w]
				wordFreq[w] = 0
			}
		}
	}

	// remove words with 0 frequency by first creating a new map
	// and add word with fequency that is not 0
	newMap := make(map[string]int)
	for k, v := range wordFreq {
		if v != 0 {
			newMap[k] = v
		}
	}
	return newMap
}
