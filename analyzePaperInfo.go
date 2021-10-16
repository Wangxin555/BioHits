package main

import (
	"regexp"
	"strings"
)

func analyzePaperInfo(papers []Paper, includeTitle bool, numHits int) []string {
	// split paragraphs into sentences
	paperSentences := make([]string, 0)
	if includeTitle {
		for i := range papers {
			processedAbstract := RemoveSpecialChar(papers[i].abstract)
			abstractSentences := strings.Split(processedAbstract, ".")
			paperSentences = append(paperSentences, abstractSentences...)
			processedTitle := RemoveSpecialChar(papers[i].title)
			paperSentences = append(paperSentences, processedTitle)
		}
	} else {
		for i := range papers {
			processedAbstract := RemoveSpecialChar(papers[i].abstract)
			abstractSentences := strings.Split(processedAbstract, ".")
			paperSentences = append(paperSentences, abstractSentences...)
		}
	}

	// paste all sentences into a long string
	// remove stop words
	senteneInOne := strings.Join(paperSentences, " ")

	return paperSentences
}

func RemoveSpecialChar(s string) string {
	tabChar := regexp.MustCompile(`\t`)
	newlineChar := regexp.MustCompile(`\n`)
	// word := "what \t is\n"
	s1 := tabChar.ReplaceAllString(s, "")
	s2 := newlineChar.ReplaceAllString(s1, "")
	return s2
}
